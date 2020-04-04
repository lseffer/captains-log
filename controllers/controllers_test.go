package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func init() {
	// Execute from project root so that relative paths are resolved correctly
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
}

func getRequest(t *testing.T, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestIndexController(t *testing.T) {
	req := getRequest(t, "/")
	fmt.Println(filepath.Abs("."))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexController)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.HasPrefix(rr.Body.String(), "<html>") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}
