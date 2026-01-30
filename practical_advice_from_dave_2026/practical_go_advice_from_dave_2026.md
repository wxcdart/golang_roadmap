# Practical Go: Real World Advice for Writing Maintainable Go Programs (2026 Update)

Based on presentations by Dave Cheney at QCon China and GopherCon Singapore 2019, updated for Go 1.24/1.25 with modern features like generics, modules, and enhanced tooling.

## Guiding Principles

Go's design prioritizes three core principles over performance or concurrency alone.

### Clarity
- Code is written for humans first, machines second. It will be read hundreds of times more than written.
- Prioritize maintainability for long-term software engineering.
- Effective communication of ideas is the most important programming skill.
- Ask yourself: "Will the next person understand my intent?"

**Example:**
```go
// Bad: Unclear intent
if x > 0 && y < 100 {
    process(x, y)
}

// Good: Clear intent
const maxValue = 100
if x > 0 && y < maxValue {
    process(x, y)
}

// Better: Self-documenting
func isValidRange(x, y int) bool {
    const maxValue = 100
    return x > 0 && y < maxValue
}

if isValidRange(x, y) {
    process(x, y)
}
```

### Simplicity
- Simplicity is essential for reliability; complexity leads to unreliable software.
- Avoid over-engineering; simple designs are harder to get right but prevent deficiencies.
- Go aims to control complexity in programming.
- When two solutions solve the same problem, choose the simpler one.
- Complexity budget: every feature added should justify its complexity cost.

**Example:**
```go
// Bad: Over-engineered with unnecessary interface
type Processor interface {
    Process() error
}

type DataProcessor struct {
    data []byte
}

func (d *DataProcessor) Process() error {
    // simple processing
    return processData(d.data)
}

// Good: Simple and direct
func processData(data []byte) error {
    // simple processing
    return nil
}
```

### Productivity
- Focus on developer efficiency: fast compilation, readable code, and minimal tooling friction.
- Go enforces consistent formatting (via `gofmt`) to reduce cognitive load and spot errors visually.
- Enables "software engineering at scale" with quick iterations and confidence in changes.

## Identifiers (Naming)

Good naming is crucial due to Go's minimal syntax; names must convey intent clearly.

### Choose Identifiers for Clarity, Not Brevity
- Optimize for readability, not line count or typing speed.
- Qualities of good names: concise, descriptive, predictable, and idiomatic.
- Describe purpose/behavior/result, not implementation or contents.

### Identifier Length
- Use the "right" length: short for local variables (e.g., `p` in loops), longer for broader scope.
- Avoid type names in variables (e.g., no `personSlice` for `[]Person`).
- Constants describe values, not usage.
- Guidelines: single letters for loops/branches, single words for params/returns, multiple for functions/packages.

**Examples by Scope:**
```go
// Loop indices: single letters
for i := 0; i < len(items); i++ {
    // ...
}

// Short-lived locals: abbreviated but clear
for _, p := range products {
    procesProduct(p)
}

// Function parameters: single word
func calculateTotal(price, quantity int) int {
    return price * quantity
}

// Package-level or long-lived: descriptive
var defaultConnectionTimeout = 30 * time.Second

// Bad: Type in name
var userSlice []User      // Don't
var userMap map[int]User  // Don't

// Good: Describe the collection
var users []User
var usersByID map[int]User
```

### Consistent Naming and Declaration Style
- Follow Go conventions (e.g., camelCase, exported names capitalized).
- Use consistent declarations (e.g., `var` vs. `:=`).
- Be a team player: match project style for collaboration.

## Comments

Comments enhance clarity without duplicating code.

### Comments on Variables/Constants
- Describe contents, not purpose (e.g., `// maxRetries is the maximum number of retries` is bad; explain what it holds).
- Good comments explain "why", not "what" (the code shows "what").

**Examples:**
```go
// Bad: Restates the obvious
// maxRetries is the maximum number of retries
const maxRetries = 3

// Good: Explains the value choice
// maxRetries is set to 3 to balance reliability with responsiveness.
// Testing showed > 3 retries rarely succeed and cause user-visible delays.
const maxRetries = 3

// Bad: Redundant
// UserCache caches users
type UserCache struct {}

// Good: Adds context
// UserCache provides thread-safe caching of user records with TTL-based eviction.
type UserCache struct {
    mu    sync.RWMutex
    cache map[int]*User
    ttl   time.Duration
}
```

### Document Public Symbols
- Always comment exported functions, types, etc.
- Start comments with the symbol name.
- Exception: Skip for interface implementations (e.g., no need for `Read` method docs if it implements `io.Reader`).

**Examples:**
```go
// ProcessPayment handles payment processing and returns a transaction ID.
// It returns an error if the payment fails validation or processing.
func ProcessPayment(amount float64, method string) (string, error) {
    // ...
}

// Config holds application configuration loaded from environment variables.
type Config struct {
    Port     int
    Database string
}
```

## Package Design

Packages are Go's unit of modularity; design them thoughtfully.

### Good Package Names
- Start with a descriptive, unique name (e.g., `http` not `web`).
- Avoid generic names like `util`, `common`, `base`.
- Package name is part of the API (e.g., `bufio.Reader`).
- Use singular nouns (e.g., `user`, not `users`).
- Short, single-word names preferred.

**Examples:**
```go
// Bad package names
package utils      // Too generic
package helpers    // Vague
package myproject  // Not descriptive

// Good package names
package user       // Clear domain
package payment    // Specific responsibility
package cache      // Describes functionality
```

### Best Practices

**Return Early to Avoid Deep Nesting:**
```go
// Bad: Deep nesting
func process(user *User) error {
    if user != nil {
        if user.IsActive {
            if user.HasPermission("write") {
                // do work
                return nil
            } else {
                return ErrNoPermission
            }
        } else {
            return ErrInactive
        }
    } else {
        return ErrNilUser
    }
}

// Good: Early returns, flat structure
func process(user *User) error {
    if user == nil {
        return ErrNilUser
    }
    if !user.IsActive {
        return ErrInactive
    }
    if !user.HasPermission("write") {
        return ErrNoPermission
    }
    
    // do work
    return nil
}
```

**Make Zero Values Useful:**
```go
// Good: Zero value is ready to use
type Buffer struct {
    data []byte  // nil slice is valid
}

func (b *Buffer) Write(p []byte) (int, error) {
    b.data = append(b.data, p...)  // Works even if b.data is nil
    return len(p), nil
}

// Usage: no initialization needed
var buf Buffer
buf.Write([]byte("hello"))  // Just works!
```

**Avoid Package-Level State:**
```go
// Bad: Global state
var cache = make(map[string]string)

func Get(key string) string {
    return cache[key]  // Not thread-safe, hard to test
}

// Good: Dependency injection
type Cache struct {
    mu   sync.RWMutex
    data map[string]string
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]string),
    }
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.data[key]
}
```

## Project Structure
- Prefer fewer, larger packages over many small ones for cohesion.
- Keep `package main` minimal; delegate to libraries.
- Leverage `go.work` (workspace mode, Go 1.18+) for multi-module projects to simplify development.

## API Design

### Make APIs Hard to Misuse
```go
// Bad: Can create invalid state
type Server struct {
    Address string
    Port    int
}

func (s *Server) Start() error {
    if s.Address == "" || s.Port == 0 {
        return errors.New("invalid config")
    }
    // ...
}

// Good: Constructor ensures validity
type Server struct {
    address string  // unexported
    port    int
}

func NewServer(address string, port int) (*Server, error) {
    if address == "" || port == 0 {
        return nil, errors.New("invalid config")
    }
    return &Server{address: address, port: port}, nil
}
```

### Optimize for Default Use Cases
```go
// Good: Common case is simple
func NewClient(endpoint string) *Client {
    return NewClientWithOptions(endpoint, DefaultOptions)
}

func NewClientWithOptions(endpoint string, opts Options) *Client {
    // Advanced configuration
}
```

### Accept Interfaces, Return Concrete Types
```go
// Good: Flexible input, concrete output
func ProcessData(r io.Reader) (*Result, error) {
    // Can accept any reader
    data, err := io.ReadAll(r)
    if err != nil {
        return nil, err
    }
    return &Result{Data: data}, nil
}
```

### Handle nil Gracefully
```go
// Good: Nil-safe methods
type User struct {
    Name string
    Age  int
}

func (u *User) String() string {
    if u == nil {
        return "<nil>"
    }
    return u.Name
}
```

### Generics Best Practices (Go 1.18+)
```go
// Good: Generic when it adds value
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// Good: Constraints for type safety
type Number interface {
    ~int | ~int64 | ~float64
}

func Sum[T Number](vals []T) T {
    var total T
    for _, v := range vals {
        total += v
    }
    return total
}

// Bad: Over-genericized
func Process[T any](t T) T {  // Just use interface{} or specific type
    return t
}
```

## Error Handling

Errors are values; treat them as such for robustness.

### Key Rules and Examples

**1. Errors Are Opaque**
```go
// Bad: Inspecting error strings
if err != nil && strings.Contains(err.Error(), "not found") {
    // Fragile!
}

// Good: Use error wrapping and Is/As (Go 1.13+)
var ErrNotFound = errors.New("not found")

if errors.Is(err, ErrNotFound) {
    // Handle not found
}
```

**2. Assert Behavior, Not Type**
```go
// Good: Check capabilities
type temporary interface {
    Temporary() bool
}

if te, ok := err.(temporary); ok && te.Temporary() {
    // Retry transient errors
}
```

**3. Never Use nil to Signal Failure**
```go
// Bad: nil has dual meaning
func FindUser(id int) *User {
    // nil could mean "not found" or "error"
    return nil
}

// Good: Explicit error
func FindUser(id int) (*User, error) {
    // ...
    if notFound {
        return nil, ErrUserNotFound
    }
    return user, nil
}
```

**4. Don't Panic in Libraries**
```go
// Bad: Panic in library code
func MustConnect(url string) *Connection {
    conn, err := Connect(url)
    if err != nil {
        panic(err)  // Harsh for library users
    }
    return conn
}

// Good: Return errors
func Connect(url string) (*Connection, error) {
    // ...
    return conn, err
}

// OK: Must* functions in main for convenience
func main() {
    conn := MustConnect(url)  // Fine in main
    defer conn.Close()
    // ...
}
```

**5. Eliminate Errors by Design**
```go
// Bad: Error-prone API
func Write(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    // ...
}

// Good: API prevents invalid state
func Write(data []byte) error {
    // Accept empty slices gracefully
    if len(data) == 0 {
        return nil  // No-op, not an error
    }
    // ...
}
```

**6. Handle Errors Once**
```go
// Bad: Error handled multiple times
func process() error {
    err := doWork()
    if err != nil {
        log.Printf("error: %v", err)  // Logged here
        return fmt.Errorf("work failed: %w", err)  // And propagated
    }
    return nil
}

func main() {
    err := process()
    if err != nil {
        log.Printf("error: %v", err)  // Logged again!
    }
}

// Good: Handle at appropriate level
func process() error {
    err := doWork()
    if err != nil {
        return fmt.Errorf("work failed: %w", err)  // Just wrap
    }
    return nil
}

func main() {
    err := process()
    if err != nil {
        log.Printf("fatal error: %v", err)  // Handle once
        os.Exit(1)
    }
}
```

**7. Use Structured Logging (Go 1.21+)**
```go
import "log/slog"

func processRequest(id string) error {
    err := validate(id)
    if err != nil {
        slog.Error("validation failed",
            "id", id,
            "error", err,
        )
        return err
    }
    return nil
}
```

## Testing

Tests ensure maintainability and enable fearless changes.

### Core Practices

**1. Table-Driven Tests**
```go
func TestCalculateDiscount(t *testing.T) {
    tests := []struct {
        name     string
        price    float64
        discount float64
        want     float64
    }{
        {"no discount", 100, 0, 100},
        {"10% off", 100, 0.1, 90},
        {"full discount", 100, 1.0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := CalculateDiscount(tt.price, tt.discount)
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

**2. Test Helpers**
```go
// Good: Helper function
func assertNoError(t *testing.T, err error) {
    t.Helper()  // Marks as helper for better error reporting
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
}

func TestProcess(t *testing.T) {
    err := process()
    assertNoError(t, err)
}
```

**3. Internal vs External Tests**
```go
// Internal test (same package): can test unexported items
// file: user_test.go
package user

func TestInternalFunction(t *testing.T) {
    result := internalFunction()  // Can access unexported
    // ...
}

// External test (different package): tests public API only
// file: user_external_test.go
package user_test

import "myapp/user"

func TestPublicAPI(t *testing.T) {
    u := user.New("alice")  // Only exported items
    // ...
}
```

**4. Fuzzing for Edge Cases (Go 1.18+)**
```go
func FuzzParseEmail(f *testing.F) {
    // Seed corpus
    f.Add("user@example.com")
    f.Add("invalid")
    
    f.Fuzz(func(t *testing.T, input string) {
        email, err := ParseEmail(input)
        if err != nil {
            return  // Invalid input is OK
        }
        
        // Valid emails must round-trip
        if email.String() != input {
            t.Errorf("failed to round-trip: %q", input)
        }
    })
}

// Run: go test -fuzz=FuzzParseEmail
```

**5. Testable Design**
```go
// Bad: Hard to test (depends on real clock)
func IsBusinessHours() bool {
    hour := time.Now().Hour()
    return hour >= 9 && hour < 17
}

// Good: Testable (inject dependencies)
type Clock interface {
    Now() time.Time
}

func IsBusinessHours(clock Clock) bool {
    hour := clock.Now().Hour()
    return hour >= 9 && hour < 17
}

// Test with fake clock
type FakeClock struct {
    current time.Time
}

func (f *FakeClock) Now() time.Time {
    return f.current
}
```

## Concurrency

Go's concurrency model requires discipline.

### Channel Axioms and Patterns

**1. Channels Are for Communication**
```go
// Bad: Using channels for synchronization
func worker(done chan bool) {
    // do work
    done <- true  // Just for signaling
}

// Better: Use sync.WaitGroup for synchronization
func worker(wg *sync.WaitGroup) {
    defer wg.Done()
    // do work
}
```

**2. Buffered Channel Sizing**
```go
// Unbuffered: Synchronous communication
ch := make(chan int)  // Size 0

// Size 1: Non-blocking send for common case
ch := make(chan int, 1)

// Avoid large buffers: they hide backpressure issues
ch := make(chan int, 1000)  // Why 1000? Usually a smell
```

**3. Keep Busy or Do the Work**
```go
// Bad: Spawning goroutine for trivial work
func process(data []byte) {
    go func() {
        result := expensiveWork(data)
        // now what? Hard to get result back
    }()
}

// Good: Let caller decide concurrency
func process(data []byte) Result {
    return expensiveWork(data)
}

// Caller can use concurrency if needed
go func() {
    result := process(data)
    // handle result
}()
```

**4. Always Know When Goroutines Stop**
```go
// Bad: Goroutine leak
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // Blocks forever if nobody sends
        process(val)
    }()
    // Goroutine never exits!
}

// Good: Use context for cancellation
func noLeak(ctx context.Context) {
    ch := make(chan int)
    go func() {
        select {
        case val := <-ch:
            process(val)
        case <-ctx.Done():
            return  // Goroutine can exit
        }
    }()
}
```

**5. Context for Cancellation (Standard Pattern)**
```go
import "context"

func processWithTimeout(ctx context.Context, data []byte) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    resultCh := make(chan error, 1)
    go func() {
        resultCh <- doWork(data)
    }()
    
    select {
    case err := <-resultCh:
        return err
    case <-ctx.Done():
        return ctx.Err()  // timeout or cancelled
    }
}
```

**6. errgroup for Goroutine Management**
```go
import "golang.org/x/sync/errgroup"

func processBatch(ctx context.Context, items []Item) error {
    g, ctx := errgroup.WithContext(ctx)
    
    for _, item := range items {
        item := item  // Capture loop variable
        g.Go(func() error {
            return processItem(ctx, item)
        })
    }
    
    return g.Wait()  // Waits for all goroutines, returns first error
}
```

**7. Worker Pool Pattern**
```go
func workerPool(ctx context.Context, jobs <-chan Job) error {
    const numWorkers = 5
    g, ctx := errgroup.WithContext(ctx)
    
    for i := 0; i < numWorkers; i++ {
        g.Go(func() error {
            for {
                select {
                case job, ok := <-jobs:
                    if !ok {
                        return nil  // Channel closed
                    }
                    if err := job.Process(); err != nil {
                        return err
                    }
                case <-ctx.Done():
                    return ctx.Err()
                }
            }
        })
    }
    
    return g.Wait()
}
```

**8. Type-Safe Channels with Generics (Go 1.18+)**
```go
// Generic pipeline stage
func pipeline[T, U any](ctx context.Context, in <-chan T, fn func(T) U) <-chan U {
    out := make(chan U)
    go func() {
        defer close(out)
        for {
            select {
            case v, ok := <-in:
                if !ok {
                    return
                }
                select {
                case out <- fn(v):
                case <-ctx.Done():
                    return
                }
            case <-ctx.Done():
                return
            }
        }
    }()
    return out
}
```

## Modern Tooling and Productivity (2026 Best Practices)

### Development Tools

**Formatting and Linting:**
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

**Managing Dev Tools (Go 1.24+):**
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

### Performance and Profiling

**Profiling Best Practices:**
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

**Profile-Guided Optimization (PGO) - Go 1.20+:**
```bash
# 1. Build with profiling
go build -o myapp

# 2. Run production workload
./myapp  # Collects CPU profile to default.pgo

# 3. Rebuild with PGO
go build -pgo=default.pgo -o myapp

# Typical improvements: 2-14% performance gain
```

### Security Best Practices (2026)

**1. Vulnerability Scanning:**
```bash
# Check for known vulnerabilities
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...

# In CI/CD pipeline
govulncheck -json ./... > vulns.json
```

**2. Dependency Management:**
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

**3. Secure Coding Patterns:**
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

### Observability and Debugging (2026)

**Structured Logging with slog (Go 1.21+):**
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

**OpenTelemetry Integration:**
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

**Metrics Collection:**
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

### Modern CI/CD Pipeline (2026)

**GitHub Actions Example:**
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

### Module Management Best Practices

**Semantic Versioning:**
```bash
# Release a new version
git tag v1.2.3
git push origin v1.2.3

# Major version bumps require new module path
# go.mod for v2+
module example.com/myapp/v2

go 1.25
```

**Workspace Mode for Multi-Module Projects (Go 1.18+):**
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

**Private Modules:**
```bash
# Configure for private repos
go env -w GOPRIVATE=github.com/myorg/*

# Use netrc for authentication
# ~/.netrc
machine github.com
    login <token>
    password x-oauth-basic
```

### Container Best Practices (2026)

**Multi-Stage Docker Build:**
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

**Distroless Images (Smaller, More Secure):**
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

### Performance Optimization Patterns (2026)

**1. Use sync.Pool for Frequently Allocated Objects:**
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

**2. Preallocate Slices When Size is Known:**
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

**3. Use strings.Builder for String Concatenation:**
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

**4. Avoid Unnecessary Conversions:**
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

### Code Organization for Scale (2026)

**Hexagonal/Clean Architecture:**
```
myapp/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── domain/           # Business logic
│   │   ├── user.go
│   │   └── order.go
│   ├── ports/            # Interfaces
│   │   ├── repositories.go
│   │   └── services.go
│   ├── adapters/         # Implementations
│   │   ├── postgres/
│   │   ├── redis/
│   │   └── http/
│   └── config/
├── pkg/                  # Public libraries
│   └── middleware/
├── go.mod
└── go.sum
```

**Feature-Based Organization (Alternative):**
```
myapp/
├── cmd/
├── internal/
│   ├── user/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── model.go
│   ├── order/
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   └── shared/
│       ├── database/
│       └── middleware/
└── go.mod
```

### Go 1.24/1.25 Specific Features

**1. Improved Type Inference for Generics:**
```go
// Go 1.25: Better inference, less need to specify types
m := Map(items, func(i Item) string {
    return i.Name
})
// No need for Map[Item, string]
```

**2. Enhanced Error Handling (Ongoing Proposals):**
```go
// Keep watching for error handling improvements
// Current best practice: wrap with context
if err != nil {
    return fmt.Errorf("failed to process user %s: %w", userID, err)
}
```

**3. Faster Builds with Improved Caching:**
```bash
# Build cache is now more aggressive
# Use GODEBUG for cache debugging
GODEBUG=gocacheverify=1 go build

# Clean cache if needed
go clean -cache -modcache
```

### Summary: 2026 Go Development Checklist

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