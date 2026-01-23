# httptest examples

This folder demonstrates using the `net/http/httptest` package to test HTTP handlers and clients.

Contents:

- `server.go` — small HTTP handlers and a function to build the `http.Handler` for tests.
- `server_test.go` — examples using `httptest.NewRecorder` (unit testing handlers) and `httptest.NewServer` (integration-style tests without real network dependencies).

Run tests:

```bash
cd golang_roadmap/04_Tooling_testing_and_code_quality/06_httptest
go test -v
```

Notes:

- Use `httptest.NewRecorder` to test handlers directly (fast, no network).
- Use `httptest.NewServer` to test client code or full-server interactions.
- Tests should avoid relying on external services; prefer injecting dependencies via interfaces.
