package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"context"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize Logrus logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	databaseURL := os.Getenv("DATABASE_URL")
	logger.Debugf("connecting to %s", databaseURL)
	// Connect to PostgreSQL database using environment variables
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to PostgreSQL:", err)
		return
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		logger.Fatal("Failed to ping to PostgreSQL:", err)
		return
	}

	// Initialize Redis client if REDIS_URL is provided
	redisURL := os.Getenv("REDIS_URL")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // No password
		DB:       0,  // Default DB
	})

	// Test the Redis connection
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		logger.Fatal("Failed to connect to Redis:", err)
		return
	}

	// Define a request handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		logger.Info(fmt.Sprintf("Received request from %s for %s", r.RemoteAddr, r.URL.Path))

		// Write the HTTP response

		fmt.Fprintf(w, "Hello, World!")

		// Optionally use Redis
		if redisClient != nil {
			ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
			defer cancel()
			err := redisClient.Set(ctx, "key", "value", 0).Err()
			if err != nil {
				logger.Warn("Failed to set data in Redis:", err)
			}
		}
	}

	// Register the request handler for the "/" path
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	logger.Info("HTTP server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Server error: ", err)
	}
}
