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

func signinHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/signin.html"))
    tmpl.Execute(w, nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    loginSignupButtons := template.Must(template.ParseFiles("templates/login-signup-buttons.html"))
    loginSignupButtons.Execute(w, nil)
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
    signUpForm := template.Must(template.ParseFiles("templates/signup.html"))
    signUpForm.Execute(w, nil)
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/signin", signinHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/signup-form", signUpHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
