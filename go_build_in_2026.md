# How Go Build Works (2026)

## The Toolchain

The Go toolchain consists of three unified tools, configured by `GOOS` and `GOARCH`:

| Tool | Command | Purpose |
|------|---------|---------|
| Compiler | `go tool compile` | Compiles `.go` → `.o` |
| Assembler | `go tool asm` | Assembles `.s` → `.o` |
| Linker | `go tool link` | Links `.o`/`.a` → executable |

Cross-compilation is trivial:
```bash
GOOS=linux GOARCH=arm64 go build
```

## The Build Process

```
.go files ──► compile ──► .a archive ──► link ──► executable
.s files  ──► asm ────┘
```

**For packages**: compile + pack into `.a` archive, then discard (or install to cache).

**For commands** (`package main`): compile + pack + link into binary.

## See What's Happening

```bash
go build -x ./...      # Print all commands
go build -n ./...      # Print commands without executing
go build -json ./...   # Structured JSON output (Go 1.24+)
```

## Build Caching

All compilation results are cached in `$GOCACHE` (~/.cache/go-build). Go 1.24+ also caches executables from `go run` and `go tool`.

```bash
go clean -cache        # Clear the build cache
go env GOCACHE         # Show cache location
```

## Tool Dependencies (Go 1.24+)

Track project tools in `go.mod`:
```
module example.com/myapp

go 1.25

tool (
    golang.org/x/tools/cmd/stringer
    honnef.co/go/tools/cmd/staticcheck
)
```

```bash
go get -tool golang.org/x/tools/cmd/stringer   # Add a tool
go tool stringer ./...                          # Run it
```

## Key Flags

| Flag | Purpose |
|------|---------|
| `-o <file>` | Output file name |
| `-v` | Print package names as compiled |
| `-x` | Print commands being run |
| `-race` | Enable race detector |
| `-trimpath` | Remove file paths for reproducible builds |
| `-ldflags` | Pass flags to linker |
| `-buildvcs=false` | Omit version control info |

## Build vs Install

- **`go build`**: Compiles; result in current dir (commands) or discarded (packages)
- **`go install`**: Compiles and installs to `$GOBIN` or `$GOPATH/bin`

## Debugging Toolchain Selection

```bash
GODEBUG=toolchaintrace=1 go build ./...
```

## Platform Requirements

**Go 1.25:**
- Linux: Kernel 3.2+
- macOS: 12 Monterey+
- Windows: 10+

**Go 1.24:**
- Linux: Kernel 3.2+
- macOS: 11 Big Sur+ (last version to support it)
- Windows: 10+

## Go 1.24 Caveats

### TLS Post-Quantum Timeouts
The new X25519MLKEM768 key exchange is enabled by default. Some buggy TLS servers can't handle the larger records, causing handshake timeouts.

```bash
GODEBUG=tlsmlkem=0 ./myapp   # Disable post-quantum key exchange
```

### Swiss Tables Map Implementation
Go 1.24 uses a new map implementation. If you see odd behavior:

```bash
GOEXPERIMENT=nosynchashtriemap go build ./...   # Revert to old maps
```

### Update to Latest Patch
Use **Go 1.24.11** — earlier versions have security fixes for net/http, os, crypto/x509, and other packages.

### `go tool` Directive
The new `tool` directive in `go.mod` may have issues with some generators (e.g., early gqlgen versions). If `go generate` fails silently, try the old `tools.go` pattern.

### Experimental: `testing/synctest`
Requires `GOEXPERIMENT=synctest` at build time. API may change in Go 1.25+.

## How Go Uses Go to Build Itself

Go's self-hosting nature means the Go toolchain is built using Go itself, starting from a minimal C-based bootstrap to avoid circular dependencies. This process ensures portability and minimal external dependencies. The build from source follows a staged approach, evolving from the original 2013 description but maintaining core principles.

### The Bootstrap Process (Adapted for Go 1.25)

Building Go from source involves several steps, typically orchestrated by scripts in `$GOROOT/src`. All commands are run from this directory.

1. **Environment Setup and Validation**  
   Run `make.bash` (or `all.bash` for full build + tests). This performs sanity checks, detects host OS/arch, and compiles `cmd/dist` (a C program) using GCC.

2. **cmd/dist (Core Build Tool)**  
   `cmd/dist` handles platform detection, code generation in `pkg/runtime`, and builds the initial compilers. It creates a minimal `go_bootstrap` toolchain, stubbing out dependencies like `cgo` to avoid circularity.

3. **go_bootstrap (Minimal Go Toolchain)**  
   Use `go_bootstrap` to compile the full standard library and replace the `go` tool. This builds a complete toolchain for the host platform.

4. **Full Toolchain Build**  
   With the bootstrap complete, build the full Go ecosystem, including compilers, linkers, and tools. Cross-compilation is enabled by default.

5. **Testing and Validation**  
   Run `run.bash` to execute tests: standard library unit tests (`go test std`), runtime/CGO tests, compiler/runtime tests (`go run test`), and API compatibility checks (`go tool api`).

### Key Differences in 2026

- **Modules Integration**: Modern builds incorporate `go.mod` for dependency management, with module proxy support for reproducible builds.
- **Enhanced Cross-Compilation**: Tools like `goreleaser` or built-in `go build` flags simplify multi-platform builds.
- **Security and Reproducibility**: Builds may include SBOM generation, vulnerability scanning, and `-trimpath` for deterministic outputs.
- **CI/CD Automation**: The process is often automated in pipelines (e.g., GitHub Actions), with caching for faster iterations.

For detailed source build instructions, see the [Go source installation guide](https://go.dev/doc/install/source).