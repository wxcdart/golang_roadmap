# Package Design

[← Back to Index](README.md) | [Next: Project Structure →](05_project_structure.md)

Packages are Go's unit of modularity; design them thoughtfully.

## Good Package Names

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

## Best Practices

### Return Early to Avoid Deep Nesting

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

### Make Zero Values Useful

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

### Avoid Package-Level State

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

---

[← Previous: Comments and Documentation](03_comments_documentation.md) | [Next: Project Structure →](05_project_structure.md)
