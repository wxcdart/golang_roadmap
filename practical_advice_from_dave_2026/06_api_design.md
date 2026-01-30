# API Design

[← Back to Index](README.md) | [Next: Error Handling →](07_error_handling.md)

## Make APIs Hard to Misuse

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

## Optimize for Default Use Cases

```go
// Good: Common case is simple
func NewClient(endpoint string) *Client {
    return NewClientWithOptions(endpoint, DefaultOptions)
}

func NewClientWithOptions(endpoint string, opts Options) *Client {
    // Advanced configuration
}
```

## Accept Interfaces, Return Concrete Types

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

## Handle nil Gracefully

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

## Generics Best Practices (Go 1.18+)

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

## API Design Checklist

- ✅ Prevent invalid states through constructors
- ✅ Make common cases simple, advanced cases possible
- ✅ Accept interfaces for flexibility
- ✅ Return concrete types for clarity
- ✅ Handle nil gracefully in methods
- ✅ Use generics only when they add real value
- ✅ Document expected behavior and edge cases

---

[← Previous: Project Structure](05_project_structure.md) | [Next: Error Handling →](07_error_handling.md)
