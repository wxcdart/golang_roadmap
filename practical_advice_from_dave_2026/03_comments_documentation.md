# Comments and Documentation

[← Back to Index](README.md) | [Next: Package Design →](04_package_design.md)

Comments enhance clarity without duplicating code.

## Comments on Variables/Constants

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

## Document Public Symbols

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

---

[← Previous: Identifiers and Naming](02_identifiers_naming.md) | [Next: Package Design →](04_package_design.md)
