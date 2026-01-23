// hello_world.go: Hello program that accepts command-line flags.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// The short variable declaration `:=` both declares and initializes a
	// new variable with its type inferred from the right-hand side. Here
	// `flag.String` returns *string, so `name` is a `*string`.
	// Use `*name` to dereference the pointer and obtain the actual string
	// value; `*shout` dereferences the `*bool` returned by `flag.Bool`.
	// Example: fmt.Println(*name) prints the name value. You can also
	// assign through the pointer: `*name = "Alice"` (pointer must be non-nil).
	name := flag.String("name", "World", "name to greet")
	shout := flag.Bool("shout", false, "print greeting in uppercase")
	flag.Parse()

	if *name == "" {
		fmt.Fprintln(os.Stderr, "error: name cannot be empty")
		os.Exit(2)
	}

	greeting := fmt.Sprintf("Hello, %s!", *name)
	if *shout {
		greeting = strings.ToUpper(greeting)
	}

	fmt.Println(greeting)
}
