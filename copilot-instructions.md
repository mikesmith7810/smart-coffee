---
name: "smart-coffee"
description: "Smart Coffee API — Enterprise Go HTTP service with structured routing, domain models, and handlers. Use when: working on handlers, routers, domain models, or adding new API endpoints."
---

# Smart Coffee API — Copilot Guidelines

## Code Style & Formatting

- **Format on save**: All Go code must be formatted with `gofmt` before commits
- **Lint**: Run `golangci-lint run ./...` to catch style violations
- **Line length**: Keep lines under 120 characters for readability
- **Comments**: Export all public functions and types with clear GoDoc comments

```go
// GetCoffee retrieves a coffee beverage by ID.
// It validates the required ID parameter and returns a Coffee domain object.
func GetCoffee(c *gin.Context) {
    // implementation
}
```

## Enterprise HTTP API Structure

- **Layered architecture**: Separate concerns into `domain`, `handlers`, `router`, and `services`
  - `domain/`: Business models and entities (structs, logic)
  - `handlers/`: HTTP request/response handling (Gin context)
  - `router/`: Route registration and middleware setup
  - `services/`: Business logic and data layer (future: database, caching)

- **Handler patterns**:
  - All handlers must accept `*gin.Context` as parameter
  - Validate input parameters early; return `400 Bad Request` if validation fails
  - Return structured JSON responses with `c.JSON(statusCode, payload)`
  - Use appropriate HTTP status codes (200, 201, 400, 404, 500)

- **Error handling**:
  - Define errors in a dedicated `errors.go` or use domain-scoped errors
  - Return errors as JSON with a consistent format: `{"error": "message"}`
  - Log errors with context (request ID, path, method)

- **Domain models**:
  - Keep domain structs in `domain/` package, independent of HTTP frameworks
  - Use JSON struct tags with kebab-case for API contracts: `json:"field_name"`
  - Include validation logic on domain types when appropriate

## Project Layout

```
smart-coffee/app/
├── main.go              # Entry point, server startup
├── go.mod              # Module definition
├── go.sum              # Dependency hashes
├── domain/             # Business entities
│   └── coffee.go       # Coffee model
├── handlers/           # HTTP handlers
│   └── coffee.go       # Coffee endpoint
├── router/             # Route registration
│   └── router.go       # Router setup
└── services/           # (Future) Business logic layer
```

## When Adding New Endpoints

1. Define domain models in `domain/`
2. Create handler in `handlers/` with validation
3. Register route in `router/New()`
4. Test with `go test ./...`
5. Format with `gofmt` before submitting

## Dependencies

- **Gin**: Web framework
- Standard library: `net/http`, `encoding/json`, `log`
- No external dependencies beyond Gin (prefer stdlib when possible)

## Testing & Validation

- Write tests in `*_test.go` files alongside implementation
- Use table-driven tests for multiple scenarios
- Aim for >80% code coverage on handlers and domain logic

## Documentation & Maintenance

**Keep copilot-instructions.md and README.md up to date with all significant changes:**
- New endpoints added? Update README API section
- Architecture changes? Update both README and instructions
- New dependencies or tools? Update prerequisites
- Breaking changes to request/response formats? Document immediately
- New test patterns or conventions? Add to instructions

**This ensures:**
- New developers onboarding understand current best practices
- Copilot gives accurate guidance on project conventions
- README remains the single source of truth for running/using the service
