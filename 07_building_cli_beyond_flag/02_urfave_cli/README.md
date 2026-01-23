urfave/cli example

This folder contains a small example demonstrating `urfave/cli` (v3) for building command-line applications.

The example shows a root command with a `--name` flag and a `greet` subcommand that takes a positional argument.

Quick start:

```bash
cd C:/Users/chang/dev/antigravity/misc/golang_roadmap/07_building_cli_beyond_flag/02_urfave_cli
go mod tidy
go run main.go --name Alice
go run main.go greet Bob
```

Features shown:
- Flags and aliases
- Subcommands
- Argument parsing (using a simple helper for flag access in this example)

Resources:
- https://github.com/urfave/cli
- https://zetcode.com/golang/urfave-cli/
