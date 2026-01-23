# Modules & Dependency Management

This folder demonstrates common module workflows and commands:

- `go mod init <module-path>` — initialize a module
- `go mod tidy` — add/remove dependencies to match imports
- Semantic import versions: use `/v2` in module path for v2+ when publishing

Example commands:

```bash
# initialize the example module (local example)
go mod init example.com/getting-started

# tidy dependencies after adding imports
go mod tidy

# build/run the module
go run .
```

See `go.mod` and `main.go` in this folder for a minimal module example.
