package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	requestCount int64
	startTime    = time.Now()
)

func main() {
	// Health check endpoint
	http.HandleFunc("/health", healthHandler)

	// Metrics endpoint
	http.HandleFunc("/metrics", metricsHandler)

	// Main endpoint
	http.HandleFunc("/", mainHandler)

	log.Println("🚀 Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&requestCount, 1)
	fmt.Fprintf(w, "OK\n")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&requestCount, 1)
	uptime := time.Since(startTime).Seconds()
	fmt.Fprintf(w, "requests_total %d\nuptime_seconds %.1f\n", atomic.LoadInt64(&requestCount), uptime)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt64(&requestCount, 1)
	fmt.Fprintf(w, "Concurrent Backend Server v1.0\n")
}
