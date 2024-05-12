package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

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
