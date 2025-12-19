// Login form
const loginForm = document.getElementById("loginForm");
const loginMessage = document.getElementById("loginMessage");
if (loginForm) {
  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formData = new FormData(loginForm);
    const res = await fetch("/login", {
      method: "POST",
      body: formData
    });
    const data = await res.json();
    loginMessage.textContent = data.message;
    if (data.success) {
      window.location.href = "/"; // Redirect to dashboard
    }
  });
}

// Signup form
const signupForm = document.getElementById("signupForm");
const signupMessage = document.getElementById("signupMessage");
if (signupForm) {
  signupForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formData = new FormData(signupForm);
    if (formData.get("password") !== formData.get("confirmPassword")) {
      signupMessage.textContent = "Passwords do not match";
      return;
    }

    const res = await fetch("/signup", {
      method: "POST",
      body: formData
    });
    const data = await res.json();
    signupMessage.textContent = data.message;
    if (data.success) {
      window.location.href = "/login"; // Redirect to login
    }
  });
}
