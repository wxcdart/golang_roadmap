// logging_slog.go
// Demonstrates the log/slog package (Go 1.21+) following Dave Cheney's logging philosophy.
//
// Based on Dave Cheney's article: "Let's Talk About Logging" (https://dave.cheney.net/2015/11/05/lets-talk-about-logging)
//
// Key Principles:
// - Only two log levels are necessary: Debug and Info
// - Debug: for developers during development or troubleshooting
// - Info: for operational awareness and user-facing messages
// - Avoid Warning level: it's either informational or an error
// - Avoid Fatal level: return errors to caller; handle cleanup in main()
// - Avoid Error level: either handle errors or return them; don't log at error level
//
// When an error occurs:
// 1. If you HANDLE it -> log at Info level
// 2. If you RETURN it -> don't log it; let the caller decide
// 3. At the boundary (main) -> you may exit after logging
//
// This example shows:
// - Structured logging with only Debug and Info levels
// - Proper error handling following Cheney's philosophy
// - Output to stdout and optional file (LOG_TO_FILE)
// - Configurable debug mode (via LOG_DEBUG)

package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// Logger with minimal levels following Cheney's philosophy:
// Only Debug and Info are necessary.
var (
	logger *slog.Logger
	debug  bool // Controlled by LOG_DEBUG environment variable
)

const (
	levelDebug = iota
	levelInfo
)

func main() {
	// Configure debug mode from environment variable
	// LOG_DEBUG=1 to enable debug logging, otherwise only info is shown
	if v := os.Getenv("LOG_DEBUG"); strings.ToLower(strings.TrimSpace(v)) == "1" || v == "true" {
		debug = true
	}

	// Build writer: stdout by default, optionally also a file
	var w io.Writer = os.Stdout
	if p := os.Getenv("LOG_TO_FILE"); p != "" {
		f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open log file %s: %v\n", p, err)
		} else {
			// Write to both stdout and file
			w = io.MultiWriter(os.Stdout, f)
			// Note: you may want to keep the file handle open for the program lifetime
			// and close it on shutdown. For rotation, use a library such as
			// github.com/natefinch/lumberjack
		}
	}

	// Use a TextHandler for human-readable output; JSONHandler is also available
	logger = slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: slog.LevelDebug, // Let the logger emit all levels; we filter manually
	}))

	// ===== Example 1: Simple info logging for operational awareness =====
	logInfo("application.start", "version", "0.1.0", "environment", "development")
	logInfo("database.connect", "host", "localhost", "port", 5432)

	// ===== Example 2: Debug logging for developers (enabled with LOG_DEBUG=1) =====
	logDebug("config.loaded", "file", "/etc/app.conf", "size_bytes", 1024)

	// ===== Example 3: Handling an error (log at Info, not Error) =====
	// When you HANDLE an error, log at Info level to indicate what you did about it
	if err := tryPlanA(); err != nil {
		logInfo("plan_a_failed_trying_plan_b", "error", err.Error())
		planB()
	}

	// ===== Example 4: Returning an error (don't log it) =====
	// When you RETURN an error, let the caller decide how to handle it
	if err := criticalOperation(); err != nil {
		// Don't log here; just return it
		logInfo("critical_operation_failed_returning_to_caller", "error", err.Error())
		os.Exit(1) // In main, we can exit directly after logging
	}

	// ===== Example 5: Demonstrating the antipattern to avoid =====
	// This is what Dave Cheney warns against:
	exampleAntipattern()

	logInfo("application.shutdown", "reason", "main.main completed successfully")
}

// ===== Helper functions following Cheney's two-level philosophy =====

// logDebug logs a message at Debug level.
// Debug logs are only emitted when LOG_DEBUG=1 is set.
// Use this for messages that help developers understand code behavior.
func logDebug(msg string, kv ...any) {
	if !debug {
		return
	}
	logger.Debug(msg, kv...)
}

// logInfo logs a message at Info level.
// Info logs are always emitted.
// Use this for operational awareness and messages users should know about.
func logInfo(msg string, kv ...any) {
	logger.Info(msg, kv...)
}

// ===== Example functions demonstrating error handling patterns =====

// tryPlanA demonstrates an operation that might fail.
// It returns an error for the caller to handle.
func tryPlanA() error {
	// Simulate a failure
	return errors.New("plan A could not open the required file")
}

// planB demonstrates a fallback plan.
func planB() {
	logInfo("plan_b.running", "status", "executing fallback logic")
}

// criticalOperation demonstrates an operation that fails and we want to exit.
func criticalOperation() error {
	// Simulate a critical failure
	return errors.New("critical operation failed: database unreachable")
}

// exampleAntipattern shows what NOT to do, with explanation.
func exampleAntipattern() {
	logInfo("=== Example: The Antipattern (to avoid) ===")

	// WRONG: Logging at Error level and returning the error
	// This violates Cheney's principle: you've already handled it by logging,
	// so why return the error?
	if err := os.Remove("some_file.txt"); err != nil {
		// This is the antipattern:
		// logger.Error("failed to remove file", "err", err)  // DON'T do this
		// return err  // Contradictory: are you handling it or returning it?

		// CORRECT: Either handle it (log at Info) or return it (don't log)
		logInfo("failed_to_remove_temporary_file_continuing", "err", err.Error())
		// ... continue execution, problem is handled
	}

	logInfo("=== Antipattern Example Ended ===")
}
