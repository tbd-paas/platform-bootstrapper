package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Channel to receive log messages
var logCh = make(chan string, 100)

// Function to handle HTTP requests
func handleLogs(w http.ResponseWriter, r *http.Request) {
	// Set headers for Server Sent Events (SSE)
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Start infinite loop to stream logs
	for {
		// Check if the client's connection is still open
		if f, ok := w.(http.Flusher); ok {
			select {
			case logMsg := <-logCh:
				// Send log message to client
				fmt.Fprintf(w, "data: %s\n\n", logMsg)
				f.Flush()
			case <-r.Context().Done():
				// Client disconnected, stop streaming
				return
			}
		} else {
			// Flusher not supported, log error
			log.Println("Error: Flusher not supported")
			return
		}
		// Introduce slight delay to avoid high CPU usage
		time.Sleep(100 * time.Millisecond)
	}
}

// Function to generate log messages
func generateLogs() {
	// Simulate generating log messages
	for i := 0; ; i++ {
		logCh <- fmt.Sprintf("Log message %d", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Start generating log messages in a separate goroutine
	go generateLogs()

	// Register handler for /logs endpoint
	http.HandleFunc("/logs", handleLogs)

	// Start HTTP server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
