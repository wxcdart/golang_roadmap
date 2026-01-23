# Essential Tooling

This folder contains notes and a tiny example illustrating essential Go tooling commands:

- `go run` — build & run a program in one step
- `go build` — compile binaries
- `go test` — run tests
- `go fmt` / `gofmt` — format code
- `go vet` — static checks for suspicious constructs

Quick uses:

```bash
# run the example
go run hello_tooling.go

# format the file
go fmt hello_tooling.go

# vet the file
go vet hello_tooling.go

# build a binary
go build -o hello_tooling.exe hello_tooling.go
```

The example `hello_tooling.go` demonstrates a tiny program you can run and test these tools against.
