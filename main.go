package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type PageData struct {
    Message string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
	  data := PageData{Message: "Henlo, VVorld!"}
    tmpl.Execute(w, data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    // This is a dummy login handler, you can replace it with your logic
    // For demonstration purposes, we just return a message
    if r.Method == http.MethodGet {
        w.Write([]byte(`<button hx-post="/logout" hx-trigger="click" hx-swap="outerHTML" hx-target="#nav-login" hx-indicator=".loading">Logout</button>`))
    } else if r.Method == http.MethodPost {
        // Logout logic (not implemented for demonstration)
        w.Write([]byte(`<button hx-get="/login" hx-trigger="click" hx-swap="outerHTML" hx-target="#nav-login" hx-indicator=".loading">Login</button>`))
    }
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
    // This is a dummy update handler, you can replace it with your logic
    // For demonstration purposes, we just return a new message
    message := "Hello, World! (Updated)"
    w.Write([]byte(message))
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/logout", loginHandler)
    http.HandleFunc("/update", updateHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
