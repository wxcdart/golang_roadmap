package main

import (
	"errors"
	"fmt"
)

// Demonstrates explicit error handling, custom error types, wrapping, and inspection.

var ErrNotFound = errors.New("not found")
var ErrNetwork = errors.New("network error")

type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string { return fmt.Sprintf("MyError %d: %s", e.Code, e.Msg) }

// get simulates operations that may return different kinds of errors.
func get(mode int) (string, error) {
	switch mode {
	case 0:
		return "payload", nil
	case 1:
		// sentinel error
		return "", ErrNotFound
	case 2:
		// custom error type
		return "", MyError{Code: 42, Msg: "custom failure"}
	case 3:
		// wrapped error using %w
		return "", fmt.Errorf("temporary network failure: %w", ErrNetwork)
	default:
		return "", fmt.Errorf("unknown mode %d", mode)
	}
}

// fetch adds context before propagating the error using %w.
func fetch(name string, mode int) (string, error) {
	res, err := get(mode)
	if err != nil {
		return "", fmt.Errorf("fetch %q failed: %w", name, err)
	}
	return res, nil
}

// higher-level function that handles errors (example of where to inspect or propagate)
func runDemo(mode int) {
	fmt.Println("--- mode", mode, "---")
	_, err := fetch("resource", mode)
	if err != nil {
		// explicit handling
		fmt.Println("error (wrapped):", err)

		// errors.Is checks for sentinel errors in the chain
		if errors.Is(err, ErrNotFound) {
			fmt.Println("-> detected: ErrNotFound")
		}
		if errors.Is(err, ErrNetwork) {
			fmt.Println("-> detected: ErrNetwork")
		}

		// errors.As extracts a concrete error type from the chain
		var me MyError
		if errors.As(err, &me) {
			fmt.Printf("-> extracted MyError: code=%d msg=%s\n", me.Code, me.Msg)
		}

		// otherwise, propagate or handle at this level
		return
	}
	fmt.Println("success")
}

func main() {
	// Demonstrate several modes
	for mode := 0; mode <= 4; mode++ {
		runDemo(mode)
		fmt.Println()
	}
	fmt.Println("Tip: always add context before returning errors and inspect them at the right level.")
}
