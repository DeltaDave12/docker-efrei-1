package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new Logrus logger instance
	logger := logrus.New()

	// Set the log level to Info
	logger.SetLevel(logrus.InfoLevel)

	// Define a request handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		logger.Info(fmt.Sprintf("Received request from %s for %s", r.RemoteAddr, r.URL.Path))

		// Write the HTTP response
		fmt.Fprintf(w, "Hello, World!")
	}

	// Register the request handler for the "/" path
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	logger.Info("HTTP server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Server error: ", err)
	}
}
