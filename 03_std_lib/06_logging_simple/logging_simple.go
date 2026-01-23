// logging_simple.go
// Demonstrates basic usage of the Go log/slog package (Go 1.21+).
//
// Note: log/slog is part of the Go standard library starting from Go 1.21.
// If you are using an earlier version, slog will not be available in the stdlib.
//
// This example shows:
// - Structured logging with log levels
// - Logging with fields
// - Logging errors
// - Customizing log output

package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// This example follows best-practices from community guides:
// - Structured logs with key/value fields
// - Configurable minimum log level (via LOG_LEVEL)
// - Output to stdout and optional file (LOG_TO_FILE)
// - Notes about rotation (use a rotation library such as lumberjack)

var (
	logger   *slog.Logger
	minLevel = levelInfo
)

const (
	levelDebug = iota
	levelInfo
	levelWarn
	levelError
)

func parseLevel(s string) int {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return levelDebug
	case "warn", "warning":
		return levelWarn
	case "error":
		return levelError
	default:
		return levelInfo
	}
}

func shouldLog(lvl int) bool { return lvl >= minLevel }

func main() {
	// Configure min level from env
	if v := os.Getenv("LOG_LEVEL"); v != "" {
		minLevel = parseLevel(v)
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
	logger = slog.New(slog.NewTextHandler(w, nil))

	// Example logs with structured fields
	info("service.start", "version", "0.1.0")
	name := "Alice"
	info("user.greet", "user", name)

	// simulate an error and log it
	err := os.Remove("nonexistent_file.txt")
	if err != nil {
		errLog(err, "file.remove.failed", "path", "nonexistent_file.txt")
	}

	warn("cache.miss", "key", "session:12345")
	debug("worker.step", "step", 3)

	customErr := errors.New("something went wrong")
	errLog(customErr, "operation.failed", "op", "doThing")

	// Fatal example (log then exit)
	// errLog(errors.New("fatal problem"), "fatal", "reason", "demo")
	// os.Exit(1)
}

// helper wrappers that respect minLevel and add consistent structure
func debug(msg string, kv ...any) {
	if !shouldLog(levelDebug) {
		return
	}
	logger.Debug(msg, kv...)
}

func info(msg string, kv ...any) {
	if !shouldLog(levelInfo) {
		return
	}
	logger.Info(msg, kv...)
}

func warn(msg string, kv ...any) {
	if !shouldLog(levelWarn) {
		return
	}
	logger.Warn(msg, kv...)
}

func errLog(err error, msg string, kv ...any) {
	if !shouldLog(levelError) {
		return
	}
	// attach the error under the key "err" for consistency
	kv2 := append([]any{"err", err}, kv...)
	logger.Error(msg, kv2...)
}
