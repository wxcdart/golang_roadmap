# Go Project Setup Guide (2026)

## Quick Start

**The modern way:** Use Go modules (the standard since Go 1.11, now mandatory).

```bash
# Create your project directory
mkdir myproject
cd myproject

# Initialize Go module
go mod init github.com/yourusername/myproject
```

This creates `go.mod` - that's it! No more `$GOPATH` gymnastics.

## Project Structures by Type

### 1. Simple CLI Tool or Small Library

**Start here.** Don't over-engineer.

```
myproject/
├── go.mod
├── go.sum
├── main.go           # or myproject.go for a library
└── main_test.go
```

### 2. Growing Project (add as needed)

When you need internal packages:

```
myproject/
├── go.mod
├── go.sum
├── main.go
└── internal/         # Private code, not importable by other projects
    ├── auth/
    │   ├── auth.go
    │   └── auth_test.go
    └── storage/
        ├── storage.go
        └── storage_test.go
```

### 3. Multiple Commands

```
myproject/
├── go.mod
├── cmd/
│   ├── server/       # Build with: go build ./cmd/server
│   │   └── main.go
│   └── worker/       # Build with: go build ./cmd/worker
│       └── main.go
└── internal/
    └── shared/
```

### 4. Server/API Project (the full structure)

Only use when your project has gotten complex:

```
myproject/
├── go.mod
├── cmd/
│   └── api-server/
│       └── main.go
├── internal/              # Private packages
│   ├── auth/
│   ├── handler/          # HTTP handlers
│   ├── service/          # Business logic
│   └── repository/       # Data access
├── pkg/                  # Public, reusable packages (optional)
│   └── validator/
├── api/                  # API specs (OpenAPI, protobuf)
├── configs/              # Config files
├── migrations/           # Database migrations
└── scripts/              # Build/deploy scripts
```

## Key Directories Explained

- **`internal/`** - Code that can't be imported by other projects. Use this for most of your code!
- **`cmd/`** - Entry points for multiple binaries. Each subdirectory becomes a separate executable.
- **`pkg/`** - Optional. Use only if you're exposing stable, reusable packages to other projects.
- **`api/`** - API definitions (OpenAPI specs, protobuf files)

## Essential Commands

```bash
# Add a dependency (automatically updates go.mod)
go get github.com/some/package

# Install specific version
go get github.com/some/package@v1.2.3

# Update all dependencies
go get -u ./...

# Remove unused dependencies
go mod tidy

# Verify checksums
go mod verify

# Build
go build ./cmd/myapp

# Run tests
go test ./...

# Run current directory
go run .
```

## Working with Local Modules (Development)

When developing multiple modules locally that depend on each other:

```bash
# In your consumer module, redirect to local path
go mod edit -replace example.com/greetings=../greetings

# Then sync dependencies
go mod tidy
```

This adds a `replace` directive to your `go.mod`:

```go
module example.com/hello

go 1.21

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```

**For production:** Remove the `replace` directive and use actual version numbers:
```go
require example.com/greetings v1.1.0
```

## Best Practices for 2026

1. **Start simple** - Begin with flat structure, add complexity only when needed
2. **Use `internal/`** - Keeps your internal code private and refactorable
3. **Avoid generic names** - No `utils`, `helpers`, `common`. Use specific names like `validator`, `parser`
4. **Keep it flat** - Go prefers flat hierarchies. Avoid deep nesting
5. **One `go.mod` per repo** - Most projects need only one module
6. **Commit `go.sum`** - Ensures reproducible builds

## Common Mistakes to Avoid

❌ Creating deep package hierarchies  
✅ Keep packages flat and focused

❌ Using `pkg/` for everything  
✅ Use `internal/` for private code, `pkg/` only for stable public APIs

❌ Generic package names (`utils`, `models`)  
✅ Descriptive names (`validator`, `user`, `payment`)

## What Changed Since 2014?

- **No more `$GOPATH` required** - Go modules handle everything
- **No more `vendor/` needed** - Go modules cache dependencies properly
- **`internal/` is now standard** - Enforced by the toolchain
- **`pkg/` is controversial** - Many prefer keeping everything in root or `internal/`

The official Go team recommends: start simple, add structure only when complexity demands it.

## Contributing to Open Source (2026 vs Old Articles)

**Old way (pre-modules, articles from 2014-2017):**
- Clone projects into `$GOPATH/src/github.com/user/project`
- Complex git remote juggling to get imports working
- Fork had to live at original import path

**Modern way (with Go modules):**
```bash
# Fork the project on GitHub, then:
git clone https://github.com/YOUR-USERNAME/project.git
cd project
go mod download  # Get dependencies

# Make changes, run tests
go test ./...

# Commit and push to your fork
git push origin your-branch

# Create PR on GitHub - done!
```

Go modules mean your fork can live anywhere - the import paths in `go.mod` handle everything. No more `$GOPATH` gymnastics!
