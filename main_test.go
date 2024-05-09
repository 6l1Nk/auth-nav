package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestDatabaseSchema(t *testing.T) {
	initDatabase()

	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		t.Fatalf("error opening datavase %v", err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT name
		FROM sqlite_master
		WHERE type='table' AND name='users'
	`)
	if err != nil {
		t.Fatalf("error querying database: %v", err)
	}
	defer rows.Close()

	if !rows.Next() {
		t.Error("users table not found in the database")
	}

	columns, err := db.Query("PRAGMA table_info(users)")
	if err != nil {
		t.Fatalf("error querying database: %v", err)
	}
	defer columns.Close()

	emailFound := false
	passwordFound := false
	for columns.Next() {
		var (
			id         int
			name       string
			dataType   string
			allowNull  int
			defaultVal interface{}
			primaryKey int
		)
		if err := columns.Scan(&id, &name, &dataType, &allowNull, &defaultVal, &primaryKey); err != nil {
			t.Fatalf("error scanning columns: %v", err)
		}
		if name == "email" {
			emailFound = true
		}
		if name == "password" {
			passwordFound = true
		}
	}

	if !emailFound {
		t.Error("email column not found in the users table")
	}

	if !passwordFound {
		t.Error("password column not found in the users table")
	}
}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "https://127.0.0.1:8443/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v and expected %v",
			status,
			http.StatusOK)
	}

	if !strings.Contains(res.Body.String(), "<nav>") {
		t.Errorf("handler returned unexpected body: missing <nav>")
	}

	loginButtonHTML := "<button id=\"login\">"
	if !strings.Contains(res.Body.String(), loginButtonHTML) {
		t.Errorf("handler returned unexpected body: missing %v", loginButtonHTML)
	}

	signUpButtonHTML := "<button id=\"signup\">"
	if !strings.Contains(res.Body.String(), signUpButtonHTML) {
		t.Errorf("handler returned unexpected body: missing %v", signUpButtonHTML)
	}
}

func TestSignUpHandler_Get(t *testing.T) {
	req, err := http.NewRequest("GET", "https://127.0.0.1:8443/sign-up", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(signUpHandler)

	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf(
			"handler returned wrong status code: got %v and expected %v",
			status,
			http.StatusOK)
	}

	signUpFormHTML := "<form action=\"/signup\" method=\"post\" class=\"signup-form\">"

	if !strings.Contains(res.Body.String(), signUpFormHTML) {
		t.Errorf("handler returned unexpected body: missing %v",
			signUpFormHTML)
	}
}

func TestSignUpHandler_Post(t *testing.T) {
	formData := strings.NewReader("email=test@example.com&password=123456")
	req, err := http.NewRequest("POST", "/sign-up", formData)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(signUpHandler)
	handler.ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
