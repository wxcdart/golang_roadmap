# Error Naming in Go
## Applying Kevlin Henney's "Exceptional Naming" Principles — Updated for 2026

*Inspired by Kevlin Henney's article "Exceptional Naming" (2021), adapted for Go's error-as-values paradigm*

---

## Introduction

Kevlin Henney's influential article "Exceptional Naming" challenged the automatic suffixing of exception classes with `Exception` in Java and .NET. His core argument: the suffix is redundant because the syntactic context (`throw`, `catch`, `throws`) already signals that something is an exception.

Go takes a radically different approach—errors are values, not exceptions—but Henney's underlying principles about naming, redundancy, and the DRY principle apply directly. In fact, Go's conventions demonstrate both alignment with and departure from his recommendations in interesting ways.

---

## Go's Error Naming Conventions

Go has established clear conventions for error naming, codified in style guides from Google, Uber, and enforced by linters like `errname`.

### The Two-Convention System

| What You're Naming | Convention | Example |
|-------------------|------------|---------|
| Sentinel error variables | Prefix with `Err` (exported) or `err` (unexported) | `ErrNotFound`, `errTimeout` |
| Custom error types | Suffix with `Error` | `NotFoundError`, `ValidationError` |

This creates an interesting asymmetry that Henney might find curious: **variables get a prefix, types get a suffix**.

---

## Sentinel Errors: The `Err` Prefix

Sentinel errors are predeclared error values used for comparison:

```go
// Standard library examples
var ErrNotExist = errors.New("file does not exist")
var ErrPermission = errors.New("permission denied")
var ErrClosed = errors.New("use of closed connection")

// Your code
var ErrNotFound = errors.New("item not found")
var ErrInvalidInput = errors.New("invalid input")
var ErrRateLimited = errors.New("rate limit exceeded")
```

### Why the `Err` Prefix Works in Go

Unlike Java's `Exception` suffix, Go's `Err` prefix serves a **distinct purpose**:

1. **No syntactic context**: In Go, errors aren't syntactically privileged. You can't tell from `if x == SomeValue` whether `SomeValue` is an error or any other comparable value. The `Err` prefix signals intent.

2. **Namespace disambiguation**: `NotFound` could be a boolean, a function, or a type. `ErrNotFound` is unambiguously an error.

3. **Discoverability**: In IDE autocomplete, typing `Err` surfaces all error values in a package.

```go
// Without prefix - ambiguous
var NotFound = errors.New("not found")  // Is this an error? A status? A flag?

// With prefix - clear intent
var ErrNotFound = errors.New("not found")  // Unambiguously an error
```

### Applying Henney's Test

Henney's test: drop the suffix and see if the name still communicates. For Go sentinel errors, try dropping the `Err` prefix:

| With Prefix | Without Prefix | Still Clear? |
|-------------|----------------|--------------|
| `ErrNotFound` | `NotFound` | ❌ Ambiguous |
| `ErrTimeout` | `Timeout` | ❌ Could be a duration |
| `ErrClosed` | `Closed` | ❌ Could be a boolean |
| `ErrPermissionDenied` | `PermissionDenied` | ⚠️ Better, but still unclear |
| `ErrInvalidFormat` | `InvalidFormat` | ⚠️ Somewhat clear |

**Conclusion**: Unlike Java's redundant `Exception` suffix, Go's `Err` prefix adds genuine semantic value because errors lack syntactic privilege.

---

## Custom Error Types: The `Error` Suffix

For struct types that implement the `error` interface, Go convention uses the `Error` suffix:

```go
// Standard library examples
type PathError struct {
    Op   string
    Path string
    Err  error
}

type SyntaxError struct {
    Offset int64
    msg    string
}

// Your code
type ValidationError struct {
    Field   string
    Message string
}

type NotFoundError struct {
    Resource string
    ID       string
}
```

### The Henney Critique Applied

This is where Henney's argument gains traction. Let's apply his test:

| With Suffix | Without Suffix | Still Clear? |
|-------------|----------------|--------------|
| `PathError` | `PathProblem` | ✅ Yes |
| `SyntaxError` | `SyntaxFault` | ✅ Yes |
| `ValidationError` | `InvalidInput` | ✅ Yes |
| `NotFoundError` | `NotFound` | ⚠️ Needs context |
| `TimeoutError` | `Timeout` | ❌ Ambiguous |

Some types genuinely benefit from the `Error` suffix for disambiguation, while others might communicate better with more descriptive names.

### Better Names Following Henney's Principles

```go
// Convention-following (suffix)
type ConnectionError struct { ... }
type ParseError struct { ... }

// Henney-inspired (descriptive)
type ConnectionRefused struct { ... }
type MalformedInput struct { ... }
type ResourceExhausted struct { ... }
type DeadlineExceeded struct { ... }
```

The descriptive approach communicates the *specific condition*, not just the category.

---

## The 2026 Consensus: Pragmatic Naming

After years of community practice, Go has settled on pragmatic conventions that balance Henney's principles with practical needs.

### The Standard Pattern

```go
package mypackage

import "errors"

// Sentinel errors: Err prefix
var (
    ErrNotFound       = errors.New("mypackage: not found")
    ErrInvalidInput   = errors.New("mypackage: invalid input")
    ErrAlreadyExists  = errors.New("mypackage: already exists")
)

// Custom error types: Error suffix
type ValidationError struct {
    Field   string
    Value   any
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

// Predicate functions for complex checks
func IsNotFound(err error) bool {
    return errors.Is(err, ErrNotFound)
}
```

### Uber Go Style Guide (2024)

Uber's widely-adopted style guide codifies:

```go
// Exported sentinel errors: Err prefix
var ErrBrokenLink = errors.New("link is broken")
var ErrCouldNotOpen = errors.New("could not open")

// Unexported sentinel errors: err prefix (no underscore)
var errNotFound = errors.New("not found")

// Exported error types: Error suffix
type NotFoundError struct {
    File string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("file %q not found", e.File)
}

// Unexported error types: also Error suffix
type notFoundError struct {
    file string
}
```

### Google Go Style Guide

Google emphasizes error message conventions over naming:

```go
// Error messages should be lowercase, no trailing punctuation
var ErrInvalidInput = errors.New("invalid input")  // Good
var ErrInvalidInput = errors.New("Invalid input.") // Bad

// Include package context in messages
var ErrNotFound = errors.New("repository: item not found")
```

---

## Naming the Condition, Not the Category

Henney's strongest point: **name the specific condition, not the general category**.

### Weak Names (Category-Based)

```go
// These tell you it's a database error, but not what went wrong
type DatabaseError struct { ... }
type NetworkError struct { ... }
type IOError struct { ... }
```

### Strong Names (Condition-Based)

```go
// These tell you exactly what happened
type ConnectionRefused struct {
    Host string
    Port int
}

type QueryTimeout struct {
    Query    string
    Duration time.Duration
}

type DiskFull struct {
    Path      string
    Required  int64
    Available int64
}

type RateLimitExceeded struct {
    Limit     int
    Window    time.Duration
    RetryAfter time.Time
}
```

### Real-World Examples from the Standard Library

The standard library demonstrates both patterns:

**Category-based (weaker)**:
```go
type SyscallError struct { ... }  // What syscall? What happened?
type AddrError struct { ... }     // What's wrong with the address?
```

**Condition-based (stronger)**:
```go
type DNSError struct {
    Err         string
    Name        string
    Server      string
    IsTimeout   bool
    IsTemporary bool
    IsNotFound  bool
}
// The fields tell you exactly what went wrong
```

---

## Assigned Error Variables

Beyond declared errors, Go has conventions for error variables in function scope:

```go
// Standard pattern
func DoSomething() error {
    result, err := someOperation()
    if err != nil {
        return fmt.Errorf("doing something: %w", err)
    }
    return nil
}

// Multiple errors (avoid shadowing)
func DoMultiple() error {
    _, err := first()
    _, err2 := second()  // err2, not err
    if err != nil {
        return err
    }
    return err2
}

// File-scope errors (more descriptive)
var setupErr error

func init() {
    setupErr = setup()
}
```

### The `errors.As` Pattern

When extracting typed errors, use single-letter names based on the type:

```go
var pErr *PathError
if errors.As(err, &pErr) {
    fmt.Println("path:", pErr.Path)
}

var vErr *ValidationError
if errors.As(err, &vErr) {
    fmt.Println("field:", vErr.Field)
}
```

---

## The `errname` Linter

The Go ecosystem has formalized these conventions through the `errname` linter:

```go
// BAD: Linter will flag these
type DecodeErr struct{}           // Should be DecodeError
var InvalidURLErr = errors.New()  // Should be ErrInvalidURL

// GOOD: Passes linter
type DecodeError struct{}
var ErrInvalidURL = errors.New("invalid URL")
```

The linter enforces:
- Sentinel error variables: `Err` or `err` prefix
- Error types: `Error` or `Errors` suffix

---

## Henney's Principles Applied to Go

### What Go Gets Right

1. **The `Err` prefix adds value**: Unlike Java's `Exception` suffix, Go's prefix disambiguates in a context where errors aren't syntactically special.

2. **Error messages matter**: Go emphasizes descriptive error messages over type names, following Henney's spirit of communicating the actual condition.

3. **Wrapping preserves context**: `fmt.Errorf("context: %w", err)` lets you add meaning without creating new types.

### Where Go Could Improve

1. **The `Error` suffix is often redundant**: When a type already describes a condition (`ConnectionRefused`), adding `Error` is noise.

2. **Category names persist**: Types like `NetworkError` and `DatabaseError` follow convention but violate Henney's principle of naming the specific condition.

---

## Practical Guidelines for 2026

### For Sentinel Errors

```go
// DO: Use Err prefix, name the condition
var ErrNotFound = errors.New("not found")
var ErrPermissionDenied = errors.New("permission denied")
var ErrRateLimitExceeded = errors.New("rate limit exceeded")

// DON'T: Redundant or vague
var ErrError = errors.New("error")           // Meaningless
var NotFoundError = errors.New("not found")  // Wrong convention
```

### For Error Types

```go
// DO: Descriptive condition names
type ConnectionRefused struct {
    Host string
    Port int
    Err  error
}

// ACCEPTABLE: Suffix with meaningful base
type ValidationError struct {
    Field   string
    Message string
}

// AVOID: Category names that hide the condition
type DatabaseError struct {
    Err error  // What actually happened?
}
```

### For Error Messages

```go
// DO: Lowercase, no punctuation, include context
return fmt.Errorf("parsing config at line %d: %w", line, err)

// DON'T: Sentence case, punctuation, vague
return fmt.Errorf("Failed to parse config.")
```

---

## Summary

Kevlin Henney's "Exceptional Naming" principles translate well to Go, with important adaptations:

| Henney's Principle | Java/C# Application | Go Application |
|-------------------|---------------------|----------------|
| Don't duplicate context | Drop `Exception` suffix | Keep `Err` prefix (adds value) |
| Name the condition | `InvalidNumberFormat` not `NumberFormatException` | `ErrInvalidFormat` or `MalformedInput` |
| Weak names are revealed | Dropping suffix shows weakness | Vague error types (`DatabaseError`) still weak |
| DRY applies to names | One source of truth | Error messages + wrapping over new types |

The Go community has found a pragmatic middle ground:
- **Prefixes for values** (`ErrNotFound`) — because errors aren't syntactically special
- **Suffixes for types** (`NotFoundError`) — convention, though Henney might argue for `NotFound`
- **Descriptive messages** — the real semantic content lives here

As Henney wrote: *"A naming convention should not be used to prop up weak names."* In Go, if your error type name needs the `Error` suffix to be understood, consider whether a more descriptive name would serve better.

---

## References

- Henney, Kevlin. "Exceptional Naming." Medium, May 2021.
- Uber Go Style Guide: Error Naming. github.com/uber-go/guide
- Google Go Style Guide: Best Practices. google.github.io/styleguide/go
- Go Blog: Working with Errors in Go 1.13. go.dev/blog/go1.13-errors
- errname linter. github.com/Antonboom/errname
- Littleroot Blog: Go Error Naming Conventions. littleroot.org

---

*Last updated: January 2026*
