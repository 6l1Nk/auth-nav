package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "strings"
)

func TestIndexHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
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

    if !strings.Contains(res.Body.String(),"<nav>") {
        t.Errorf("handler returned unexpected body: missing <nav>")
    }

    if !strings.Contains(res.Body.String(),"<button id=\"login\">") {
        t.Errorf("handler returned unexpected body: missing <button id=\"login\">")
    }

    if !strings.Contains(res.Body.String(),"<button id=\"signup\">") {
        t.Errorf("handler returned unexpected body: missing <button id=\"signup\">")
    }
}
