package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Demonstrates Dave Cheney's error handling philosophy:
// - Errors are values
// - Wrap with context using %w
// - Use errors.Is and errors.As for inspection
// - Design APIs to eliminate error handling (bufio.Scanner pattern)
// - Use helper types like errWriter for repetitive operations
// - Assert behavior, not type

var (
	ErrNotFound = errors.New("not found")
	ErrNetwork  = errors.New("network error")
	ErrTimeout  = errors.New("operation timeout")
)

// CustomError demonstrates a custom error type with Unwrap for error chains.
// Per Dave Cheney's advice, limit public error types - use them carefully.
type CustomError struct {
	Code int
	Msg  string
	Err  error // Wrapped error for context
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("CustomError %d: %s: %v", e.Code, e.Msg, e.Err)
	}
	return fmt.Sprintf("CustomError %d: %s", e.Code, e.Msg)
}

func (e *CustomError) Unwrap() error {
	return e.Err
}

// temporary interface for asserting behavior, not type (Dave's recommendation)
type temporary interface {
	Temporary() bool
}

// isTemporaryError checks if an error represents a temporary/retryable condition.
func isTemporaryError(err error) bool {
	var t temporary
	return errors.As(err, &t) && t.Temporary()
}

// get simulates operations that may return different kinds of errors.
func get(mode int) (string, error) {
	switch mode {
	case 0:
		return "payload", nil
	case 1:
		// sentinel error - avoid comparing directly; use errors.Is instead
		return "", ErrNotFound
	case 2:
		// custom error type with wrapped error
		return "", &CustomError{Code: 42, Msg: "custom failure", Err: ErrNetwork}
	case 3:
		// wrapped error using %w to preserve error chain
		return "", fmt.Errorf("temporary network failure: %w", ErrNetwork)
	case 4:
		// wrapping with custom type
		return "", fmt.Errorf("fetch from service: %w", &CustomError{Code: 500, Msg: "service unavailable"})
	default:
		return "", fmt.Errorf("unknown mode %d", mode)
	}
}

// fetch adds context before propagating the error using %w.
// This follows Dave Cheney's principle: wrap with context, let caller decide how to handle.
func fetch(name string, mode int) (string, error) {
	res, err := get(mode)
	if err != nil {
		// Add context using %w to preserve the error chain
		return "", fmt.Errorf("fetch %q (mode %d) failed: %w", name, mode, err)
	}
	return res, nil
}

// countLinesScanner demonstrates the "bufio.Scanner pattern" from Dave Cheney's philosophy.
// This API eliminates error handling by encapsulating the complexity.
func countLinesScanner(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	lines := 0
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

// errWriter is a helper type that accumulates I/O errors.
// From Dave Cheney's "Eliminate Error Handling by Eliminating Errors".
// Allows writing multiple operations without checking after each one.
type errWriter struct {
	w   io.Writer
	err error
}

func (ew *errWriter) Write(p []byte) (int, error) {
	if ew.err != nil {
		return 0, ew.err
	}
	var n int
	n, ew.err = ew.w.Write(p)
	return n, nil
}

func (ew *errWriter) Printf(format string, args ...any) (int, error) {
	if ew.err != nil {
		return 0, ew.err
	}
	return fmt.Fprintf(ew.w, format, args...)
}

func (ew *errWriter) Err() error {
	return ew.err
}

// runDemo demonstrates error inspection using errors.Is and errors.As.
// Per Dave Cheney: only handle errors once. This function logs and returns,
// letting the caller decide on further action.
func runDemo(mode int) {
	fmt.Println("--- mode", mode, "---")
	_, err := fetch("resource", mode)
	if err != nil {
		// Log the full error with context
		fmt.Printf("error (wrapped): %v\n", err)
		fmt.Println()

		// errors.Is checks for errors in the chain by identity
		if errors.Is(err, ErrNotFound) {
			fmt.Println("-> detected: ErrNotFound")
		}
		if errors.Is(err, ErrNetwork) {
			fmt.Println("-> detected: ErrNetwork")
		}

		// errors.As extracts a specific error type from the chain
		var ce *CustomError
		if errors.As(err, &ce) {
			fmt.Printf("-> extracted CustomError: code=%d msg=%s\n", ce.Code, ce.Msg)
		}

		fmt.Println()
		return
	}
	fmt.Println("success")
	fmt.Println()
}

// demoErrWriter shows how to eliminate error handling using helper types.
func demoErrWriter() {
	fmt.Println("--- errWriter demo ---")

	var buf bytes.Buffer
	ew := &errWriter{w: &buf}

	ew.Printf("HTTP/1.1 200 OK\r\n")
	ew.Printf("Content-Type: text/plain\r\n")
	ew.Printf("Content-Length: 13\r\n")
	ew.Printf("\r\n")
	ew.Write([]byte("Hello, World!"))

	if err := ew.Err(); err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}

	fmt.Printf("output:\n%s\n", buf.String())
	fmt.Println()
}

// demoLineCounter shows the bufio.Scanner pattern.
func demoLineCounter() {
	fmt.Println("--- line counter (Scanner pattern) ---")

	text := "line 1\nline 2\nline 3"
	r := strings.NewReader(text)

	lines, err := countLinesScanner(r)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("counted %d lines\n", lines)
	fmt.Println()
}

// demoErrorJoin demonstrates Go 1.20+ errors.Join for multiple errors.
func demoErrorJoin() {
	fmt.Println("--- errors.Join demo (Go 1.20+) ---")

	// Simulate multiple errors from concurrent operations
	err1 := fmt.Errorf("operation 1 failed: %w", ErrTimeout)
	err2 := fmt.Errorf("operation 2 failed: %w", ErrNetwork)

	combined := errors.Join(err1, err2)

	// errors.Is checks the entire tree
	if errors.Is(combined, ErrTimeout) {
		fmt.Println("found ErrTimeout in error chain")
	}
	if errors.Is(combined, ErrNetwork) {
		fmt.Println("found ErrNetwork in error chain")
	}

	fmt.Printf("combined error: %v\n", combined)
	fmt.Println()
}

func main() {
	fmt.Println("Dave Cheney's Error Handling Philosophy (Go 2026)")
	fmt.Println("================================================\n")

	// Demo 1: Error wrapping and inspection
	fmt.Println("Demo 1: Error Wrapping and Inspection")
	fmt.Println("-------------------------------------")
	for mode := 0; mode <= 4; mode++ {
		runDemo(mode)
	}

	// Demo 2: errWriter pattern
	demoErrWriter()

	// Demo 3: bufio.Scanner pattern (eliminate error handling)
	demoLineCounter()

	// Demo 4: errors.Join (Go 1.20+)
	demoErrorJoin()

	fmt.Println("Key Principles:")
	fmt.Println("- Errors are values; treat them accordingly")
	fmt.Println("- Wrap errors with context using fmt.Errorf(\"context: %w\", err)")
	fmt.Println("- Use errors.Is() for sentinel errors, errors.As() for type extraction")
	fmt.Println("- Handle errors only once (either log or return, not both)")
	fmt.Println("- Design APIs to eliminate error handling (bufio.Scanner pattern)")
	fmt.Println("- Use helper types to reduce error handling boilerplate")
	fmt.Println("- Assert behavior, not type, for maximum flexibility")
}
