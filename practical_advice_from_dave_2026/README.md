# Practical Go: Real World Advice for Writing Maintainable Go Programs (2026)

Based on presentations by Dave Cheney at QCon China and GopherCon Singapore 2019, updated for Go 1.24/1.25 with modern features like generics, modules, and enhanced tooling.

## Table of Contents

### Core Principles and Language Features

1. **[Guiding Principles](01_guiding_principles.md)**
   - Clarity, Simplicity, and Productivity
   - Code examples demonstrating each principle

2. **[Identifiers and Naming](02_identifiers_naming.md)**
   - Choosing clear identifiers
   - Identifier length guidelines
   - Consistent naming conventions

3. **[Comments and Documentation](03_comments_documentation.md)**
   - Effective commenting strategies
   - Documenting public symbols
   - When to comment and when not to

4. **[Package Design](04_package_design.md)**
   - Good package names
   - Best practices for package organization
   - Avoiding common pitfalls

5. **[Project Structure](05_project_structure.md)**
   - Modern Go project layouts
   - Feature-based vs. hexagonal architecture
   - Internal package organization

### API and Design Patterns

6. **[API Design](06_api_design.md)**
   - Making APIs hard to misuse
   - Handling nil gracefully
   - Using generics effectively
   - Accept interfaces, return concrete types

7. **[Error Handling](07_error_handling.md)**
   - Treating errors as values
   - Error wrapping and inspection
   - Structured logging with slog
   - Best practices for 2026

### Testing and Quality

8. **[Testing](08_testing.md)**
   - Table-driven tests
   - Test helpers and organization
   - Fuzzing (Go 1.18+)
   - Testable design patterns

### Concurrency

9. **[Concurrency](09_concurrency.md)**
   - Channel axioms and patterns
   - Context for cancellation
   - Worker pools and errgroup
   - Type-safe channels with generics

### Modern Tooling (2026)

10. **[Modern Tooling and Best Practices](10_modern_tooling.md)**
    - Development tools and workflow
    - Performance profiling and PGO
    - Security best practices
    - Observability and debugging
    - CI/CD pipelines
    - Container best practices
    - Performance optimization patterns
    - Go 1.24/1.25 specific features

## Quick Reference

**Essential 2026 Go Development Checklist:**
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

## About

This guide compiles practical advice from Dave Cheney's presentations and workshops, updated for modern Go development in 2026. The content emphasizes clarity, simplicity, and productivity—the three core principles of Go's design.

## Related Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Dave Cheney's Blog](https://dave.cheney.net/)
