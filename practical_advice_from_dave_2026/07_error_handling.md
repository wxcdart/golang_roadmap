# Error Handling

[← Back to Index](README.md) | [Next: Testing →](08_testing.md)

Errors are values; treat them as such for robustness.

## Key Rules and Examples

### 1. Errors Are Opaque

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

### 2. Assert Behavior, Not Type

```go
// Good: Check capabilities
type temporary interface {
    Temporary() bool
}

if te, ok := err.(temporary); ok && te.Temporary() {
    // Retry transient errors
}
```

### 3. Never Use nil to Signal Failure

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

### 4. Don't Panic in Libraries

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

### 5. Eliminate Errors by Design

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

### 6. Handle Errors Once

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

### 7. Use Structured Logging (Go 1.21+)

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

## Error Handling Checklist

- ✅ Use `errors.Is` and `errors.As` for error inspection
- ✅ Wrap errors with context using `fmt.Errorf` with `%w`
- ✅ Define sentinel errors as package variables
- ✅ Return errors, don't panic (except in main/tests)
- ✅ Handle errors at the appropriate level
- ✅ Use structured logging for better observability
- ✅ Design APIs that minimize error possibilities

---

[← Previous: API Design](06_api_design.md) | [Next: Testing →](08_testing.md)
