package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func saveNewUser(email string, hashedPassword []byte) error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func initDatabase() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		password TEXT
		)`)
	if err != nil {
		return err
	}
	return nil
}
