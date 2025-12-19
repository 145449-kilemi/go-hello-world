package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

/* ---------------- GLOBAL STATE ---------------- */

var db *sql.DB
var messages = []string{}
const maxMessages = 20

/* ---------------- AUTH MIDDLEWARE ---------------- */

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

/* ---------------- MAIN ---------------- */

func main() {
	router := gin.Default()

	/* ---------- DATABASE ---------- */
	var err error
	connStr := "user=postgres password=YOUR_PASSWORD dbname=godashboard sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	/* ---------- SESSIONS ---------- */
	store := cookie.NewStore([]byte("secret123"))
	router.Use(sessions.Sessions("gosession", store))

	/* ---------- STATIC & TEMPLATES ---------- */
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*")

	/* ---------- AUTH PAGES ---------- */

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	/* ---------- SIGNUP ---------- */
	router.POST("/signup", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(400, gin.H{"success": false, "message": "Fill all fields"})
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		_, err := db.Exec(
			"INSERT INTO users (username, password) VALUES ($1, $2)",
			username, string(hashedPassword),
		)

		if err != nil {
			c.JSON(400, gin.H{"success": false, "message": "Username already exists"})
			return
		}

		c.JSON(200, gin.H{"success": true, "message": "Signup successful"})
	})

	/* ---------- LOGIN ---------- */
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		var storedHash string
		err := db.QueryRow(
			"SELECT password FROM users WHERE username=$1",
			username,
		).Scan(&storedHash)

		if err != nil {
			c.JSON(401, gin.H{"success": false, "message": "Invalid username or password"})
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)) == nil {
			session := sessions.Default(c)
			session.Set("user", username)
			session.Save()

			c.JSON(200, gin.H{"success": true, "message": "Login successful"})
		} else {
			c.JSON(401, gin.H{"success": false, "message": "Invalid username or password"})
		}
	})

	/* ---------- LOGOUT ---------- */
	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/login")
	})

	/* ---------- DASHBOARD ---------- */
	router.GET("/", AuthRequired(), func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"user": user,
		})
	})

	/* ---------- APIs ---------- */

	router.GET("/api/data", AuthRequired(), func(c *gin.Context) {
		c.JSON(200, gin.H{"items": []string{"Go", "Gin", "Frontend", "Backend"}})
	})

	router.POST("/api/submit", AuthRequired(), func(c *gin.Context) {
		name := c.PostForm("name")
		c.JSON(200, gin.H{"message": "Received " + name})
	})

	/* ---------- LIVE MESSAGES (SSE) ---------- */

	router.GET("/stream", AuthRequired(), func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		lastIndex := 0
		for {
			if lastIndex < len(messages) {
				for _, msg := range messages[lastIndex:] {
					fmt.Fprintf(c.Writer, "data: %s\n\n", msg)
				}
				c.Writer.Flush()
				lastIndex = len(messages)
			}
			time.Sleep(500 * time.Millisecond)
		}
	})

	router.POST("/send", AuthRequired(), func(c *gin.Context) {
		msg := c.PostForm("message")
		if msg != "" {
			messages = append(messages, msg)
			if len(messages) > maxMessages {
				messages = messages[len(messages)-maxMessages:]
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	/* ---------- LIVE CHARTS ---------- */

	router.GET("/chart-stream-line", AuthRequired(), func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		for i := 1; i <= 20; i++ {
			data := fmt.Sprintf(`{"label":"Point %d","value":%d}`, i, (i*7)%30+5)
			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()
			time.Sleep(2 * time.Second)
		}
	})

	router.GET("/chart-stream-bar", AuthRequired(), func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		for i := 1; i <= 20; i++ {
			data := fmt.Sprintf(`{"label":"BPoint %d","value":%d}`, i, (i*3)%25+2)
			fmt.Fprintf(c.Writer, "data: %s\n\n", data)
			c.Writer.Flush()
			time.Sleep(2 * time.Second)
		}
	})

	router.Run(":9090")
}
