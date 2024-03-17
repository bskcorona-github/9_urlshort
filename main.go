package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

var urlMap = make(map[string]string)

const baseURL = "http://asdflokajsfakjsgp@agkja@psgjapojfaopfjas.com/"

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/redirect/", redirectHandler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL cannot be empty", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	urlMap[shortURL] = url

	fmt.Fprintf(w, "Shortened URL: %s%s\n", baseURL, shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/redirect/")
	longURL, exists := urlMap[shortURL]
	if !exists {
		http.Error(w, "Shortened URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}
