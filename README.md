# Go + Gin Real-Time Dashboard Toolkit

## Overview
This toolkit demonstrates a beginner-friendly Go web application using the **Gin framework**. The dashboard includes:

- Dynamic list fetched from backend API  
- Form submission with live response  
- Static and live charts using Chart.js  
- Live messages using Server-Sent Events (SSE)  
- Collapsible cards for UI organization  
- Dark/light mode toggle  

This project serves as a toolkit to help beginners get started with Go, Gin, SSE, and Chart.js.

---

## System Requirements
- **OS:** Windows, Linux, or Mac  
- **Tools:** VS Code (recommended) or any code editor  
- **Dependencies:**  
  - Go v1.21+  
  - Gin framework (`go get github.com/gin-gonic/gin`)  
- **Browser:** Chrome, Firefox, or Edge  

---

## Setup Instructions

1. Install Go 1.21+ from [https://go.dev/dl/]

2. Install Gin framework:
   Open terminal or PowerShell and run:

```bash
go get github.com/gin-gonic/gin

## Run the Dashboard

In the project folder, run:

```bash
go run main.go

go-dashboard/
├─ main.go            # Go backend server
├─ templates/
│   └─ index.html     # HTML dashboard
├─ assets/
│   ├─ css/
│   │   └─ style.css  # Styles
│   └─ js/
│       └─ main.js    # Frontend logic

## How It Works

### Backend (`main.go`)
- Serves HTML pages, API endpoints, and Server-Sent Event (SSE) streams.  
- Provides dynamic list data, form handling, and live updates for charts and messages.

### Frontend
- Fetches API data and renders it in the list, charts, and messages section.  
- Handles form submissions asynchronously without reloading the page.  
- Updates charts in real-time using SSE.  
- Provides UI features like collapsible cards and a dark/light mode toggle.

### Authentication
- Signup and login with **bcrypt password hashing** (secure)
- Session-based authentication using **Gin sessions**
- Logout functionality
- Protected routes (dashboard, APIs, SSE streams)

## Features & Expected Output

| Feature           | Description                           | Expected Output                                  |
|------------------|---------------------------------------|-------------------------------------------------|
| Dynamic List      | Fetches items from backend API        | Go, Gin, Frontend, Backend                      |
| Form Submission   | Submit a name and receive confirmation| Confirmation message displayed under form       |
| Static Chart      | Bar chart with sample data            | Color-coded bar chart rendered on page load     |
| Live Line Chart   | Updates every 2 seconds               | Chart dynamically updates with new points      |
| Live Bar Chart    | Updates every 2 seconds               | Chart dynamically updates with new bars        |
| Live Messages     | Real-time messages appear incrementally | Messages like “Message 1, 2…” appear in list |
| UI Controls       | Collapsible cards & dark/light mode  | Cards toggle content visibility; theme switches instantly |

---

## Common Issues & Fixes

| Issue                     | Fix                                                                 |
|----------------------------|--------------------------------------------------------------------|
| Port 8080 in use           | Change `router.Run(":8080")` to another port   9090                     |
| Charts not showing         | Wrap JS code in `document.addEventListener('DOMContentLoaded', () => {...})` |
| SSE data not updating      | Ensure headers: `Content-Type: text/event-stream`, `Cache-Control: no-cache` |


---

## Features

### Authentication
- Signup and login with **bcrypt password hashing** (secure)
- Session-based authentication using **Gin sessions**
- Logout functionality
- Protected routes (dashboard, APIs, SSE streams)

### Live Data
- **Live Line Chart** and **Live Bar Chart** using **Server-Sent Events (SSE)**
- Live messages stream
- Dynamic list with example items
- Form submission API (demo purposes)

### UI & UX
- Collapsible cards for sections with arrows indicating open/closed state
- Toggle Dark/Light mode
- Responsive dashboard layout
- Fully front-end and back-end integrated

---

## Technology Stack

- **Backend**: Go (Gin framework)
- **Database**: PostgreSQL (persistent user storage)
- **Frontend**: HTML, CSS, JS, Chart.js
- **Security**: bcrypt for password hashing
- **Real-Time Updates**: Server-Sent Events (SSE)
- **Sessions**: `github.com/gin-contrib/sessions` (cookie-based)

---

## Database Setup

1. Install PostgreSQL
2. Create database:

```sql
CREATE DATABASE godashboard;
\c godashboard

Create users table:

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

---

## AI Prompt Journal (Learning Reflections)

Add a Test User

Important: Passwords must be hashed with bcrypt. Plain text passwords will not work.

Option 1: Create via Signup Page

Run the Go server:

go run main.go

Open browser: http://localhost:9090/signup

Create a user (e.g., testuser / testpass)

Login using the credentials.

Option 2: Manual Insert with Bcrypt

Generate a bcrypt hash in Go:

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "testpass"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println(string(hash))
}


Copy the printed hash and insert into the database:

INSERT INTO users(username, password) VALUES ('testuser', '<PASTE_HASH_HERE>');


Replace <PASTE_HASH_HERE> with the generated hash.

Login will now work with username: testuser and password: testpass.

5. Configure Go Project

Update PostgreSQL connection string in main.go:

connStr := "user=postgres password=YOUR_PASSWORD dbname=godashboard sslmode=disable"


Replace YOUR_PASSWORD with your PostgreSQL password.Mine in this case is 5432

6. Run the Project
go run main.go


Server runs on http://localhost:9090.


| Prompt                                           | Reflection                                       |
|-------------------------------------------------|-------------------------------------------------|
| “Guide me to create a Go Gin dashboard with live charts” | Scaffolded backend and SSE endpoints           |
| “Collapsible cards in HTML/CSS”                | Implemented toggle buttons for better UI interaction |
| “Fix Chart.js not showing in Go project”       | Solved rendering issue by waiting for DOM load |

## References

- [Go Official Documentation](https://go.dev/doc/)  
- [Gin Framework Documentation](https://gin-gonic.com/docs/)  
- [Chart.js Documentation](https://www.chartjs.org/docs/latest/)  
