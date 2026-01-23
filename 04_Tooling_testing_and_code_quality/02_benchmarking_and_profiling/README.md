# Benchmarking & Profiling

This folder contains simple benchmarks and an example pprof server.

Commands:

```bash
cd golang_roadmap/04_Tooling_testing_and_code_quality/02_benchmarking_and_profiling
go test -bench .

# run the pprof server and visit http://localhost:6060/debug/pprof/
go run pprof_server.go
```

Notes:
- Use `go test -bench .` for microbenchmarks.
- Use `pprof` (or `go tool pprof`) to inspect CPU/memory profiles.
