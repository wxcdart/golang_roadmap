package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// Demonstrates the Go standard library's flag package.
// Shows basic flags (string, int, bool, duration), parsing, usage, subcommands,
// custom flag types, and how to access values and handle errors.

// CSVFlag is a simple custom flag.Value that parses comma-separated values.
type CSVFlag []string

func (c *CSVFlag) String() string {
	return strings.Join(*c, ",")
}

func (c *CSVFlag) Set(s string) error {
	if s == "" {
		return errors.New("empty csv value")
	}
	parts := strings.Split(s, ",")
	for _, p := range parts {
		*c = append(*c, strings.TrimSpace(p))
	}
	return nil
}

func main() {
	// Basic flags on the default FlagSet (CommandLine)
	name := flag.String("name", "world", "name to greet")
	count := flag.Int("n", 1, "number of greetings")
	verbose := flag.Bool("v", false, "enable verbose output")
	timeout := flag.Duration("timeout", 5*time.Second, "operation timeout (e.g. 500ms, 2s")

	// Custom flag type (CSV)
	var tags CSVFlag
	flag.Var(&tags, "tags", "comma-separated list of tags")

	// Customize Usage (auto-generated help still available via -h)
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "\nExamples:")
		fmt.Fprintln(flag.CommandLine.Output(), "  # basic usage")
		fmt.Fprintln(flag.CommandLine.Output(), "  go run flag_example.go -name=Alice -n=3")
		fmt.Fprintln(flag.CommandLine.Output(), "  go run flag_example.go subcmd -port=8080")
	}

	// If no subcommand, parse default flags and run main behavior
	if len(os.Args) > 1 && os.Args[1] == "subcmd" {
		// Demonstrate subcommands using a new FlagSet
		sub := flag.NewFlagSet("subcmd", flag.ExitOnError)
		port := sub.Int("port", 8080, "port to listen on for subcmd")
		debug := sub.Bool("debug", false, "enable debug for subcmd")
		// Parse the subcommand flags from os.Args[2:]
		if err := sub.Parse(os.Args[2:]); err != nil {
			fmt.Println("subcmd parse error:", err)
			os.Exit(2)
		}
		fmt.Println("subcmd running on port:", *port, "debug:", *debug)
		return
	}

	// Parse default flags
	flag.Parse()

	// Access leftover positional arguments after parsing flags
	extra := flag.Args()

	if *verbose {
		fmt.Println("flags: name=", *name, "count=", *count, "timeout=", *timeout, "tags=", tags)
		fmt.Println("extra args:", extra)
	}

	// Use values
	for i := 0; i < *count; i++ {
		fmt.Printf("Hello, %s!\n", *name)
	}

	// Simulate using duration flag
	fmt.Printf("(would wait up to %s for an operation)\n", timeout.String())

	// Demonstrate error handling example: validate flags
	if *count < 1 {
		fmt.Println("error: -n must be >= 1")
		flag.Usage()
		os.Exit(2)
	}

	// Show help auto-generation: run `-h` or `--help` to see it
}
