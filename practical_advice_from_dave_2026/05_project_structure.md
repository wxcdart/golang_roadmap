# Project Structure

[← Back to Index](README.md) | [Next: API Design →](06_api_design.md)

## Principles

- Prefer fewer, larger packages over many small ones for cohesion.
- Keep `package main` minimal; delegate to libraries.
- Leverage `go.work` (workspace mode, Go 1.18+) for multi-module projects to simplify development.

## Hexagonal/Clean Architecture

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

**When to use:**
- Large applications with complex business logic
- Need clear separation of concerns
- Multiple deployment targets (API, CLI, workers)
- Team collaboration with clear boundaries

## Feature-Based Organization (Alternative)

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

**When to use:**
- Medium-sized applications
- Features are relatively independent
- Easier to locate feature-specific code
- Quick development iterations

## Workspace Mode for Multi-Module Projects

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

**Benefits:**
- Develop multiple modules simultaneously
- No need for `replace` directives
- Test cross-module changes locally

---

[← Previous: Package Design](04_package_design.md) | [Next: API Design →](06_api_design.md)
