# Go Learning Roadmap - Examples

This repository contains a collection of small, runnable Go examples organized by topic, following the Go learning roadmap.

## Modules

1. **01_getting_started** - Basic setup, tooling, and simple programs
2. **02_core_language** - Core language features (variables, functions, structs, interfaces, concurrency)
3. **03_std_lib** - Standard library usage (flag, time, os/io, bufio, regex, embed, logging)
4. **04_Tooling_testing_and_code_quality** - Testing, benchmarking, linting, formatting
5. **05_logging_beyond_slog** - Logging backends (Zerolog, Zap, Logrus) with trace middleware
6. **06_db_access** - Database access with GORM
7. **07_building_cli_beyond_flag** - CLI frameworks (Bubble Tea, urfave CLI)
8. **08_web_development** - Web development with net/http
9. **09_rpc** - Remote Procedure Calls with net/rpc

## TODO

- [x] Create getting started examples
- [x] Create core language examples
- [x] Create std lib examples
- [x] Create testing/tooling examples
- [x] Create logging backend examples
- [x] Create DB access examples (GORM)
- [x] Create CLI examples (Bubble Tea, urfave CLI)
- [x] Create web development examples (net/http)
- [x] Create RPC examples (net/rpc)
- [ ] Add more web examples (e.g., gRPC, frameworks like Gin)
- [ ] Add advanced concurrency examples
- [ ] Add deployment/Docker examples
- [ ] Add OpenTelemetry tracing examples

Each module contains:
- `go.mod` - Module definition
- `main.go` - Runnable example
- `README.md` - Explanation and usage

Run examples with: `cd <module> && go run main.go`