package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// func signInHandler(w http.ResponseWriter, r *http.Request) {
//     tmpl := template.Must(template.ParseFiles("templates/signin.html"))
//     tmpl.Execute(w, nil)
// }

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		tmpl.Execute(w, nil)
	}

	//	if r.Method == http.MethodPost {
	//	    fmt.Println("post")
	//	    err := r.ParseForm()
	//	    if err != nil {
	//	        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	        return
	//	    }
	//
	//	    fmt.Println("canary")
	//	    // Parse form data
	//	    email := r.Form.Get("email")
	//	    password := r.Form.Get("password")
	//	    confirmPassword := r.Form.Get("confirm-password")
	//
	//	    fmt.Println(email)
	//	    fmt.Println(password)
	//	    fmt.Println(confirmPassword)
	//	    // // Check if passwords match
	//	    // if password != confirmPassword {
	//	    //     http.Error(w, "Passwords do not match", http.StatusBadRequest)
	//	    //     return
	//	    // }
	//	    //
	//	    // Check if email already exists in the database
	//	    // var count int
	//	    // err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	//	    // if err != nil {
	//	    //     http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	    //     return
	//	    // }
	//	    // if count > 0 {
	//	    //     http.Error(w, "Email already exists", http.StatusConflict)
	//	    //     return
	//	    // }
	//	    //
	//	    // // Insert new user record into the database
	//	    // _, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, password)
	//	    // if err != nil {
	//	    //     http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	    //     return
	//	    // }
	//	    //
	//	    // // Success response
	//	    // w.WriteHeader(http.StatusCreated)
	//	    // fmt.Fprintf(w, "Account created successfully!")
	//	    tmpl := template.Must(template.ParseFiles("templates/index.html"))
	//	    tmpl.Execute(w, nil)
	//	}
}

func initDatabase() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE,
        password TEXT
    )`)
	if err != nil {
		panic(err)
	}
}

func main() {
	initDatabase()

	http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/signin", signInHandler)
	// http.HandleFunc("/signup", signUpHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting user home directory: ", err)
	}
	certFile := homeDir + "/Dev/ssl/server.crt"
	keyFile := homeDir + "/Dev/ssl/server.key"

	fmt.Println("Server is running on https://127.0.0.1:8443")
	log.Fatal(http.ListenAndServeTLS(":8443", certFile, keyFile, nil))
}
