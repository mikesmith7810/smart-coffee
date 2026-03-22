# Smart Coffee

A Go-based HTTP server implementing a coffee ordering process. This project demonstrates building an enterprise-grade API with structured layering, domain-driven design, and is designed to run locally for **Grafana observability demonstrations**.

## Overview

**Smart Coffee** is a lightweight REST API service that handles coffee ordering operations. It's built as a local playground to showcase how to instrument and monitor Go services with Grafana, including metrics, logs, and traces.

**Use cases:**
- Local development and testing of coffee ordering workflows
- Grafana integration testing (metrics collection, dashboarding, alerting)
- Go HTTP API architecture reference (layered structure, error handling, validation)

## Project Structure

```
smart-coffee/
├── app/                    # Go service
│   ├── main.go            # Entry point, server startup
│   ├── go.mod / go.sum    # Go modules
│   ├── domain/            # Business entities
│   │   └── coffee.go      # Coffee model
│   ├── handlers/          # HTTP request handlers
│   │   └── coffee.go      # Coffee endpoints
│   └── router/            # Route registration
│       └── router.go      # Gin router setup
├── copilot-instructions.md # Coding guidelines
└── README.md              # This file
```

## Getting Started

### Prerequisites

- **Go 1.25.6+** ([Install Go](https://golang.org/dl/))
- **Gin** (automatically installed via `go mod`)

### Installation

```bash
cd smart-coffee/app
go mod download
```

### Starting the Server

From the `app/` directory:

```bash
go run .
```

The server will start on `http://localhost:8080` and output:
```
[GIN-debug] Loaded HTML Templates (0): 
[GIN-debug] POST   /coffee/               --> smart-coffee/handlers.GetCoffee (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
```

### Building a Binary

```bash
cd smart-coffee/app
go build -o server .
./server
```

## API Endpoints

### Get Coffee

**Endpoint:** `GET /coffee/`

**Query Parameters:**
- `id` (required): Unique identifier for the coffee order

**Response (200 OK):**
```json
{
  "id": "1",
  "name": "Latte",
  "calories": 150
}
```

**Error (400 Bad Request):**
```json
{
  "error": "Key: 'CoffeeQuery.Id' Error:Field validation for 'Id' failed on the 'required' tag"
}
```

**Example:**
```bash
curl http://localhost:8080/coffee/?id=123
```

## Architecture

### Layered Design

- **`domain/`**: Pure business models (structs, validation logic) — framework-agnostic
- **`handlers/`**: HTTP request/response layer — Gin-specific
- **`router/`**: Route registration and middleware setup
- **`services/`**: (Future) Business logic and data persistence

### Key Patterns

- **Request validation**: Early validation with `binding:"required"` tags
- **Error handling**: Consistent JSON error responses with HTTP status codes
- **Middleware**: Gin's built-in Logger and Recovery middleware

## Development

### Code Style

All code is formatted with `gofmt`:
```bash
gofmt -w ./...
```

### Linting

Run `golangci-lint` for code quality checks:
```bash
golangci-lint run ./...
```

### Testing

Run tests:
```bash
go test ./...
```

## Load Testing with k6

This project includes **k6 load testing** to validate API performance and behavior under load. k6 tests are written in JavaScript and simulate realistic user traffic patterns.

### Prerequisites

- **k6** ([Install k6](https://k6.io/docs/getting-started/installation/))

### Running Load Tests

Start the server:
```bash
cd smart-coffee/app
go run .
```

In another terminal, run k6 tests:
```bash
k6 run load-tests/coffee.js
```

### Load Test Scenarios

- **Smoke test**: Single request to verify endpoint health
- **Load test**: Sustained traffic (e.g., 50 virtual users for 30 seconds)
- **Stress test**: Gradually increase load to find breaking point
- **Spike test**: Sudden traffic spikes to test resilience

### Example Output

```
          /\      |‾‾| /‾‾/‾‾ / /‾‾/
         /  \     |  |/  /   / /  / 
        /    \    |     (   /  /
       /______\   |  |\  \ /  ‾‾\
       , |___| |,__ /  \ / \ |_  _/|
       `._    /._,' \    |   |___| v0.50.0

     execution: local
     script: load-tests/coffee.js
     output: -

scenarios: (100.00%) 1 scenario, 50 max VUs, 1m0s max duration (hold for 30s)

GET /coffee/?id=123:
  ✓ status is 200
  ✓ response time < 200ms

checks...................: 100% ✓ 1200   ✗ 0
```

Grafana can ingest k6 metrics to track API performance trends over time.

This service is instrumented to work with Grafana for local monitoring. Common integration points:

- **Metrics**: Expose Prometheus metrics (future: add instrumentation)
- **Logs**: Structured logging via Gin middleware
- **Traces**: Ready for OpenTelemetry integration

See `/grafana` folder (future) for Grafana dashboard configurations.

## Future Enhancements

- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] Prometheus metrics endpoint
- [ ] OpenTelemetry tracing
- [ ] Authentication & authorization
- [ ] Unit and integration tests
- [ ] Docker support
- [ ] Grafana dashboard templates
- [ ] k6 load test scenarios (smoke, load, stress, spike tests)
