package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`<button id="logout" hx-get="/logout" hx-trigger="click" hx-swap="innerHTML" hx-target="#nav-login" hx-indicator=".loading">Logout</button>`))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`<button id="login" hx-get="/login" hx-trigger="click" hx-swap="innerHTML" hx-target="#nav-login" hx-indicator=".loading">Login</button>`))
    w.Write([]byte(`<button id="signup">Sign Up</button>`))
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
