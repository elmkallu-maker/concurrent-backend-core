\# Concurrent Backend Server



A production-grade HTTP server in Go demonstrating concurrent request handling, metrics collection, and graceful shutdown patterns.



\## Features



\- \*\*Concurrent Request Handling\*\* – Built with Go's goroutines for efficient concurrent processing

\- \*\*Metrics Endpoint\*\* – Monitor request counts and server uptime in real-time

\- \*\*Health Check\*\* – Built-in `/health` endpoint for service monitoring

\- \*\*Graceful Shutdown\*\* – Clean server termination with connection handling



\## Architecture



\- `cmd/server/main.go` – HTTP server entry point with three endpoints

\- Atomic counters for thread-safe metrics collection

\- Standard library HTTP handling (no external dependencies)



\## Getting Started



```bash

go run ./cmd/server



