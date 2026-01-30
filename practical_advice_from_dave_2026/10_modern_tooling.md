# Modern Tooling and Best Practices (2026)

[← Back to Index](README.md)

## Development Tools

### Formatting and Linting

```bash
# Standard formatting (stdlib)
go fmt ./...

# Stricter formatting (third-party)
go install mvdan.cc/gofumpt@latest
gofumpt -l -w .

# Comprehensive linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run ./...
```

### Managing Dev Tools (Go 1.24+)

```go
// go.mod
module example.com/myapp

go 1.25

tool (
    mvdan.cc/gofumpt@latest
    github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    golang.org/x/vuln/cmd/govulncheck@latest
)
```

Run with: `go tool gofumpt` or `go tool golangci-lint`.

## Performance and Profiling

### Profiling Best Practices

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof

# Benchmarking with memory stats
go test -bench=. -benchmem -benchtime=5s

# Continuous profiling in production
# Use runtime/pprof or net/http/pprof endpoints
```

### Profile-Guided Optimization (PGO) - Go 1.20+

```bash
# 1. Build with profiling
go build -o myapp

# 2. Run production workload
./myapp  # Collects CPU profile to default.pgo

# 3. Rebuild with PGO
go build -pgo=default.pgo -o myapp

# Typical improvements: 2-14% performance gain
```

## Security Best Practices (2026)

### 1. Vulnerability Scanning

```bash
# Check for known vulnerabilities
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...

# In CI/CD pipeline
govulncheck -json ./... > vulns.json
```

### 2. Dependency Management

```bash
# Verify dependencies haven't been tampered with
go mod verify

# Update dependencies to patch versions only
go get -u=patch ./...

# Check for outdated dependencies
go list -u -m all

# Use minimal module graphs (Go 1.21+)
# go.mod automatically uses lazy loading
```

### 3. Secure Coding Patterns

```go
// Always validate inputs
func ProcessUser(input string) error {
    // Bad: Direct use
    query := "SELECT * FROM users WHERE name = '" + input + "'"
    
    // Good: Parameterized queries
    query := "SELECT * FROM users WHERE name = ?"
    rows, err := db.Query(query, input)
    // ...
}

// Use crypto/rand for security-sensitive operations
import "crypto/rand"

func generateToken() (string, error) {
    b := make([]byte, 32)
    _, err := rand.Read(b)  // Not math/rand!
    if err != nil {
        return "", err
    }
    return base64.URLEncoding.EncodeToString(b), nil
}
```

## Observability and Debugging (2026)

### Structured Logging with slog (Go 1.21+)

```go
import "log/slog"

func main() {
    // JSON logging for production
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
        AddSource: true,  // Include source location
    }))
    slog.SetDefault(logger)
    
    // Structured logging
    slog.Info("user logged in",
        "user_id", userID,
        "ip", remoteIP,
        "duration_ms", duration.Milliseconds(),
    )
    
    // With context
    logger.InfoContext(ctx, "processing request",
        "request_id", requestID,
        "method", r.Method,
    )
}
```

### OpenTelemetry Integration

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
)

func processOrder(ctx context.Context, order Order) error {
    tracer := otel.Tracer("order-service")
    ctx, span := tracer.Start(ctx, "process_order")
    defer span.End()
    
    span.SetAttributes(
        attribute.String("order.id", order.ID),
        attribute.Int("order.items", len(order.Items)),
    )
    
    // Process with traced context
    err := validateOrder(ctx, order)
    if err != nil {
        span.RecordError(err)
        return err
    }
    
    return nil
}
```

### Metrics Collection

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "HTTP request duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "endpoint", "status"},
    )
)

func init() {
    prometheus.MustRegister(requestDuration)
}

func handler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    defer func() {
        duration := time.Since(start).Seconds()
        requestDuration.WithLabelValues(
            r.Method,
            r.URL.Path,
            strconv.Itoa(http.StatusOK),
        ).Observe(duration)
    }()
    
    // Handle request
}
```

## Modern CI/CD Pipeline (2026)

### GitHub Actions Example

```yaml
name: Go CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - uses: actions/setup-go@v5
        with:
          go-version: '1.25'
          cache: true
      
      - name: Verify dependencies
        run: |
          go mod verify
          go mod tidy
          git diff --exit-code
      
      - name: Lint
        uses: golangci/golangci-lint-action@v4
        with:
          version: latest
      
      - name: Test
        run: |
          go test -race -coverprofile=coverage.out ./...
          go test -bench=. -benchmem ./...
      
      - name: Vulnerability check
        run: |
          go install golang.org/x/vuln/cmd/govulncheck@latest
          govulncheck ./...
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
```

## Module Management Best Practices

### Semantic Versioning

```bash
# Release a new version
git tag v1.2.3
git push origin v1.2.3

# Major version bumps require new module path
# go.mod for v2+
module example.com/myapp/v2

go 1.25
```

### Workspace Mode for Multi-Module Projects (Go 1.18+)

```bash
# Initialize workspace
go work init ./service-a ./service-b

# go.work file created
go 1.25

use (
    ./service-a
    ./service-b
)

# Now can develop both modules with local changes
```

### Private Modules

```bash
# Configure for private repos
go env -w GOPRIVATE=github.com/myorg/*

# Use netrc for authentication
# ~/.netrc
machine github.com
    login <token>
    password x-oauth-basic
```

## Container Best Practices (2026)

### Multi-Stage Docker Build

```dockerfile
# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath \
    -ldflags="-s -w" -o /app/server ./cmd/server

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/server .

# Non-root user
RUN adduser -D -u 10000 appuser
USER appuser

EXPOSE 8080
CMD ["./server"]
```

### Distroless Images (Smaller, More Secure)

```dockerfile
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o server .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /app/server /
USER nonroot:nonroot
CMD ["/server"]
```

## Performance Optimization Patterns (2026)

### 1. Use sync.Pool for Frequently Allocated Objects

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func processData(data []byte) ([]byte, error) {
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)
    
    // Use buffer
    buf.Write(data)
    return buf.Bytes(), nil
}
```

### 2. Preallocate Slices When Size is Known

```go
// Bad: Multiple allocations
var results []Result
for _, item := range items {
    results = append(results, process(item))
}

// Good: Single allocation
results := make([]Result, 0, len(items))
for _, item := range items {
    results = append(results, process(item))
}
```

### 3. Use strings.Builder for String Concatenation

```go
// Bad: Creates many intermediate strings
func buildQuery(fields []string) string {
    query := "SELECT "
    for i, field := range fields {
        if i > 0 {
            query += ", "
        }
        query += field
    }
    return query
}

// Good: Efficient string building
func buildQuery(fields []string) string {
    var b strings.Builder
    b.WriteString("SELECT ")
    for i, field := range fields {
        if i > 0 {
            b.WriteString(", ")
        }
        b.WriteString(field)
    }
    return b.String()
}
```

### 4. Avoid Unnecessary Conversions

```go
// Bad: Unnecessary []byte to string conversion
func contains(data []byte, substr string) bool {
    return strings.Contains(string(data), substr)
}

// Good: Use bytes package
func contains(data []byte, substr string) bool {
    return bytes.Contains(data, []byte(substr))
}
```

## Go 1.24/1.25 Specific Features

### 1. Improved Type Inference for Generics

```go
// Go 1.25: Better inference, less need to specify types
m := Map(items, func(i Item) string {
    return i.Name
})
// No need for Map[Item, string]
```

### 2. Enhanced Error Handling (Ongoing Proposals)

```go
// Keep watching for error handling improvements
// Current best practice: wrap with context
if err != nil {
    return fmt.Errorf("failed to process user %s: %w", userID, err)
}
```

### 3. Faster Builds with Improved Caching

```bash
# Build cache is now more aggressive
# Use GODEBUG for cache debugging
GODEBUG=gocacheverify=1 go build

# Clean cache if needed
go clean -cache -modcache
```

## Summary: 2026 Go Development Checklist

- ✅ Use Go 1.24+ with latest patch version
- ✅ Enable PGO for production builds
- ✅ Use `slog` for structured logging
- ✅ Implement OpenTelemetry for observability
- ✅ Run `govulncheck` in CI/CD
- ✅ Use `golangci-lint` with recommended settings
- ✅ Write table-driven tests with `-race` flag
- ✅ Use generics judiciously (not everywhere)
- ✅ Manage tools with `tool` directive in go.mod
- ✅ Profile before optimizing
- ✅ Use context for cancellation and timeouts
- ✅ Container images: multi-stage or distroless
- ✅ Keep dependencies up to date and verified
- ✅ Document public APIs with examples

---

[← Previous: Concurrency](09_concurrency.md) | [Back to Index](README.md)
