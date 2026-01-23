// http_server_simple.go
// Simple HTTP server with two endpoints, logging, and JSON responses.
//
// This example demonstrates:
// - Creating an HTTP server with net/http
// - Adding endpoints (/hello, /time)
// - Logging each request
// - Returning JSON responses using encoding/json

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Response defines the structure for JSON responses.
type Response struct {
	Message string `json:"message"`        // Main message
	Time    string `json:"time,omitempty"` // Optional: server time
}

// loggingMiddleware logs each incoming HTTP request before passing it to the next handler.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r) // Call the next handler
	})
}

// helloHandler handles requests to /hello and returns a JSON greeting.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{Message: "Hello, world!"}
	w.Header().Set("Content-Type", "application/json") // Set response type to JSON
	json.NewEncoder(w).Encode(resp)                    // Encode and send the response as JSON
}

// timeHandler handles requests to /time and returns the current server time in JSON.
func timeHandler(w http.ResponseWriter, r *http.Request) {
	resp := Response{
		Message: "Current server time",
		Time:    time.Now().Format(time.RFC3339), // RFC3339 is a standard time format
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()
	// Register handler functions for endpoints
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/time", timeHandler)

	// Wrap the mux with logging middleware
	loggedMux := loggingMiddleware(mux)

	log.Println("Starting server on :8080...")
	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
