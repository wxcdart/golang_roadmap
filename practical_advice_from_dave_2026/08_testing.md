# Testing

[← Back to Index](README.md) | [Next: Concurrency →](09_concurrency.md)

Tests ensure maintainability and enable fearless changes.

## Core Practices

### 1. Table-Driven Tests

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

### 2. Test Helpers

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

### 3. Internal vs External Tests

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

### 4. Fuzzing for Edge Cases (Go 1.18+)

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

### 5. Testable Design

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

## Testing Checklist

- ✅ Write table-driven tests for comprehensive coverage
- ✅ Use `t.Helper()` for test utility functions
- ✅ Prefer internal tests for implementation details
- ✅ Use external tests for public API validation
- ✅ Add fuzzing for parsers and input validation
- ✅ Design for testability (dependency injection)
- ✅ Run tests with `-race` flag to detect race conditions
- ✅ Measure coverage but don't obsess over 100%

## Useful Testing Commands

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...

# Run with coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific test
go test -run TestCalculateDiscount

# Run fuzzing
go test -fuzz=FuzzParseEmail -fuzztime=30s

# Verbose output
go test -v ./...

# Benchmarking
go test -bench=. -benchmem
```

---

[← Previous: Error Handling](07_error_handling.md) | [Next: Concurrency →](09_concurrency.md)
