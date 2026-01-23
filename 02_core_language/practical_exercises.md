# Core Language & Standard Library — Practical Exercises

This section provides hands-on exercises to deepen your understanding of Go's core language features and standard library. The exercises reference small, self-contained example programs under `golang_roadmap/02_core_language_std_lib/` so you can run and adapt them.

---

## Topics Covered
- Methods and interfaces
- Composition vs embedding
- Error handling idioms (explicit errors, wrapping, inspection)
- Defer, panic, and recover
- Concurrency primitives: goroutines, channels, `select`, sync package (`Mutex`, `WaitGroup`, `Once`), and `context` for cancellation/timeouts
- Memory model: value vs reference semantics, pointers, slices, arrays
- Standard library highlights:
	- `net/http`: HTTP client and server implementations for building web services and making HTTP requests.
	- `encoding/json`: Encoding and decoding JSON data, essential for APIs and data interchange.
	- `io`: Core interfaces and helpers for I/O primitives (Reader, Writer, Copy, etc.).
	- `fmt`: Formatted I/O for printing and scanning (e.g., Printf, Sprintf, Scanf).
	- `log` / `slog`: Structured logging and simple logging helpers.
	- `context`: Propagating cancellation, timeouts, and request-scoped values across API boundaries and goroutines.
	- `database/sql`: Generic interface for SQL databases (MySQL, Postgres, SQLite, etc.), supporting queries, transactions, and connection pooling.

---

## Practical Exercises (quick index)

Below are hands-on tasks mapped to the small example programs in this repo. Run each example with `go run` inside its folder and modify it to experiment.

1. Variable scope and program structure
	- File: [10_variable_scope/variable_scope.go](10_variable_scope/variable_scope.go)
	- Exercise: inspect package, imports, `main()`, `var` vs `:=`, and scope rules.

2. Data types and variables
	- File: [11_datatypes_and_variables/datatypes_and_variables.go](11_datatypes_and_variables/datatypes_and_variables.go)
	- Exercise: experiment with `int`/`float64`/`string`/`bool`, zero values, constants and `iota`, and formatted printing.

3. Flow control
	- File: [12_flow_control/flow_control.go](12_flow_control/flow_control.go)
	- Exercise: try `if/else`, `switch`, `for` variants, `range`, `break`/`continue` and labeled `break`.

4. Functions and parameters
	- File: [13_functions_parameters/functions_parameters.go](13_functions_parameters/functions_parameters.go)
	- Exercise: implement multiple returns, variadic functions, anonymous functions, and higher-order functions.

5. Arrays, slices, maps
	- File: [09_arrays_slices_2Dslices_maps/arrays_slices_maps.go](09_arrays_slices_2Dslices_maps/arrays_slices_maps.go)
	- Exercise: observe array value-copy semantics vs slice header sharing, `append`/`cap` behavior, and map operations.

6. Structs and methods
	- File: [14_structs_and_methods/structs_methods.go](14_structs_and_methods/structs_methods.go)
	- Exercise: define structs, embed types, add methods with value and pointer receivers, and compare structs vs maps for modeling data.

7. Interfaces
	- File: [15_go_interfaces/interfaces.go](15_go_interfaces/interfaces.go)
	- Exercise: learn implicit implementation, empty `interface{}`, type assertions, and type switches.

8. Error handling
	- File: [16_go_error_handling/error_handling.go](16_go_error_handling/error_handling.go)
	- Exercise: practice explicit `if err != nil` handling, custom error types, wrapping with `%w`, and `errors.Is`/`errors.As`.

9. Pointers and memory
	- File: [17_pointers_and_memory/pointers_memory.go](17_pointers_and_memory/pointers_memory.go)
	- Exercise: experiment with pointers, dereferencing, `new`, nil checks, and value vs reference semantics.

10. Concurrency & advanced topics (suggested)
	- Implement worker pools with channels and `context` for cancellation.
	- Explore `sync` primitives and `time` for timeouts.

Each example folder contains a small `main` program you can run directly. Try modifying the code and re-running to observe behavior.
---

## Tips
- Explore the official docs for [`net/http`](https://pkg.go.dev/net/http), [`encoding/json`](https://pkg.go.dev/encoding/json), [`context`](https://pkg.go.dev/context), and [`sync`](https://pkg.go.dev/sync).
- Experiment with error handling using `defer`, `panic`, and `recover` in small programs.
- Try to write self-contained, testable code for each exercise.

---

### How to run
Open a terminal in a specific example folder and run:

```bash
cd golang_roadmap/02_core_language_std_lib/12_flow_control
go run flow_control.go
```

### Tips for learning and verifying
- Add small unit tests (use `testing` package) for logic you change.
- Use `go vet` and `golangci-lint` (if available) to catch common issues.
- Add contextual errors with `fmt.Errorf("...: %w", err)` when returning errors.
- When experimenting with concurrency, prefer small reproducible examples and use `-race` detector:

```bash
go test -race ./...
```

*Save your solutions and revisit them as you learn more!* 
