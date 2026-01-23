# Flag package examples

This folder contains examples demonstrating the Go standard library `flag` package.

Contents
- `flag_example.go`: shows how to declare flags of several types (string, int, bool, duration), custom flag types, parsing, usage/help text, simple validation, and subcommands via `flag.NewFlagSet`.

Quick run

```bash
cd golang_roadmap/03_std_lib/02_flag
go run flag_example.go -name=Alice -n=3 -tags=go,example
```

Subcommand

```bash
go run flag_example.go subcmd -port=9090 -debug=true
```

Notes
- The `flag` package automatically generates `-h`/`-help` output showing defaults and descriptions.
- For more advanced CLI needs (subcommands, complex parsing), consider third-party packages like `spf13/cobra`.
