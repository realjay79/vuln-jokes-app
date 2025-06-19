package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // specific origin!
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)

	if r.Method == "OPTIONS" {
		return
	}
	jokeID := r.URL.Query().Get("id")

	// SSRF vulnerability: allowing external/internal URL injection
	jokeURL := fmt.Sprintf("https://official-joke-api.appspot.com/jokes/%s", jokeID)

	// Simulate SSRF by letting user control the URL path
	resp, err := http.Get(jokeURL)
	if err != nil {
		http.Error(w, "Failed to fetch joke", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// XSS vulnerability: reflect joke content back as raw HTML
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "%s", body)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/joke", jokeHandler)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
