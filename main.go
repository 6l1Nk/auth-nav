package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type SignUpFormData struct {
	Email    string
	Password string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// func signInHandler(w http.ResponseWriter, r *http.Request) {
//     tmpl := template.Must(template.ParseFiles("templates/signin.html"))
//     tmpl.Execute(w, nil)
// // Verify a password against the hashed password
// err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(formData.Password))
// if err != nil {
// 	fmt.Println("Password does not match hash:", err)
// 	return
// }
// fmt.Println("Password matches hash!")
// //
// }

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tmpl := template.Must(template.ParseFiles("templates/signup.html"))
		tmpl.Execute(w, nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		formData := SignUpFormData{
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}

		// Generate a salted hash of the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(formData.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		db, err := sql.Open("sqlite3", "./users.db")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", formData.Email, hashedPassword)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Redirect to a success page
		http.Redirect(w, r, "/success", http.StatusFound)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
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
	http.HandleFunc("/sign-up", signUpHandler)
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
