# Let's Talk About Logging

Based on the article by Dave Cheney: https://dave.cheney.net/2015/11/05/lets-talk-about-logging

---

## Introduction

Go's standard `log` package lacks leveled logging and per-package control. This has led to many third-party replacements, but Dave Cheney argues that most logging libraries are fundamentally flawed—they offer far too many features and log levels. The solution is to ruthlessly pare down unnecessary complexity.

---

## Why No Love for the Standard Log Package?

Go's `log` package doesn't have built-in log levels. You have to manually add prefixes like `debug`, `info`, `warn`, and `error` yourself. There's also no way to control these levels on a per-package basis.

### Third-Party Alternatives

**glog** (from Google):
- Info
- Warning
- Error
- Fatal (terminates the program)

**loggo** (from Juju):
- Trace
- Debug
- Info
- Warning
- Error
- Critical

While these libraries offer more features, they inherit their lineage from older logging systems like `syslog(3)`. However, Dave argues they are fundamentally **wrong** because they offer **too many features**—a bewildering array of choices that distracts programmers from thinking clearly about how to communicate with future readers of their logs.

> "All logging libraries are bad because they offer too many features."

---

## The Problem with Warning Level

**Nobody needs a warning log level.**

### Why?

1. **Nobody reads warnings** - By definition, if you're logging a warning, nothing has actually gone wrong yet. The problem is speculative, not actual.

2. **It's an admission of incorrect design** - If you're using leveled logging, why would you set the level to `warning`? You'd either set it to `info` (to see informational messages) or `error` (to see errors). Setting it to `warning` is an admission that you're probably logging errors at the warning level.

### Conclusion

Either something is an **informational message** or an **error condition**. There's no middle ground. Eliminate the warning level entirely.

---

## The Problem with Fatal Level

`log.Fatal` logs a message and then calls `os.Exit(1)`. This has serious consequences:

- **Defer statements in other goroutines don't run**
- **Buffers aren't flushed**
- **Temporary files and directories aren't removed**

In essence, `log.Fatal` is semantically equivalent to `panic`—it's just less verbose.

### The Industry Standard

It's widely accepted that **libraries should not use `panic`**. But if `log.Fatal` has the same effect, shouldn't it also be outlawed?

### The Cleanup Problem

Some argue that shutdown handlers registered with the logging system can solve cleanup issues. However, this introduces:
- **Tight coupling** between your logging system and cleanup operations
- **Separation of concerns** violations

### The Better Approach

**Don't log at fatal level.** Instead:
- Return errors to the caller
- Let errors bubble up to `main.main()`
- Handle cleanup actions there before exiting
- If needed, write directly to `os.Stderr` using `fmt.Fprintf`

---

## The Problem with Error Level

Error handling and logging are closely related, but logging at error level is often inappropriate.

### The Two Options

When a function returns an error, you have only two options:

1. **Handle the error**
2. **Return the error to your caller** (possibly gift-wrapped with additional context)

### The Core Issue

```go
err := somethingHard()
if err != nil {
    log.Error("oops, something was too hard", err)
    return err  // What is this, Java?
}
```

**The problem:** If you log an error and then return it, you've already handled it—so it's no longer appropriate to log it as an **error**.

### When to Log About Errors

You **should** log when an error occurs and you've chosen to handle it:

```go
if err := planA(); err != nil {
    log.Infof("couldn't open the foo file, continuing with plan b: %v", err)
    planB()
}
```

In this example, `log.Info` and `log.Error` serve the same purpose: informing the reader that a condition occurred. Since you've **handled** the error, it's no longer an error condition—it's just informational.

### The Conclusion on Error Logging

> "You should never be logging anything at error level because you should either handle the error or pass it back to the caller."

You're not being told "don't log errors." You're being told: log that errors **occurred**, but log them at the appropriate level. An overwhelming proportion of items logged at error level are actually just informational—they're simply related to an error.

---

## What's Left?

After removing `warning`, `fatal`, and `error` levels, what logging levels do you actually need?

### The Two Things Worth Logging

1. **Things developers care about** when developing or debugging software
2. **Things users care about** when using your software

These correspond to two levels:

- **`log.Debug`** - For developers during development or troubleshooting
- **`log.Info`** - For users and operational awareness

### log.Info

- Should write directly to log output
- Should **not** be optional or configurable
- Users should only be told things useful to them
- If an **unhandleable** error occurs, let it bubble to `main.main()` where the program terminates
- Writing `FATAL` to stderr or using `fmt.Fprintf` is not sufficient justification for a `log.Fatal` method

### log.Debug

- Entirely controlled by developers and support engineers
- Should be plentiful during development (no need for `trace` or `debug2` levels)
- Should support fine-grained control to enable/disable debug statements at package scope or finer
- Useful for development, debugging, and operational diagnostics

---

## The Principle

Logging is both **important** and **hard**.

The solution is to **deconstruct** and **ruthlessly pare down** unnecessary distraction. A minimal logging API with just two levels serves the vast majority of use cases better than a complex API with six or more levels.

---

## Summary

### Eliminate These Levels:
- **Warning** - It's either informational or an error
- **Error** - Return it or handle it; don't log it at error level
- **Fatal** - Use `os.Exit()` in `main()` instead; return errors to the caller

### Keep Only These:
- **Info** - For operational awareness and user-facing messages
- **Debug** - For development, troubleshooting, and diagnostic purposes

### Key Principle:
> "The smallest possible logging API is the best one."

Stop trying to predict what level each log message should be at. Ask yourself: "Does this message help developers debug the code, or does it help users understand what their program is doing?" The answer determines whether it's debug or info.

---

## slog and the Evolution of Go Logging

This article was written in 2015, long before `slog` was introduced. **slog** (structured logging) became part of Go's standard library in **Go 1.21**, and it represents a significant evolution in how Go approaches logging.

### What slog Brings to the Table

**slog** is a structured logging package that uses key-value pairs instead of unstructured text:

```go
log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
log.Info("user login", "user_id", 123, "ip_address", "192.168.1.1")
// Output: {"time":"...","level":"INFO","msg":"user login","user_id":123,"ip_address":"192.168.1.1"}
```

#### Key Features:
- **Pluggable handlers** - Route logs to different backends (JSON, text, custom)
- **Fine-grained control** - Adjust log levels per logger instance
- **Attributes context** - Attach contextual key-value pairs to log messages
- **Standard library** - No external dependencies; better performance than third-party options
- **Minimal API** - Much simpler than pre-1.21 third-party alternatives

### How slog Aligns with Cheney's Philosophy

1. **Simplicity** - slog is intentionally minimal compared to libraries like Logrus, Zap, or Zerolog
2. **Structured approach** - Encourages explicit, machine-readable logging rather than "clever" formatting
3. **Not trying to do everything** - No built-in profiling, metrics, or distributed tracing
4. **Explicit error handling** - Errors are still values; they're logged explicitly, not exceptions

### Where slog Still Falls Short (By Cheney's Standards)

1. **Multiple levels remain** - slog still includes `Debug`, `Info`, `Warn`, and `Error` levels
   - Cheney would argue `Warn` is unnecessary
   - The argument that errors should be returned or handled still applies

2. **Backwards compatibility constraints** - slog's design includes more levels to provide an upgrade path from the standard `log` package

### The Practical Takeaway

While slog doesn't fully embrace Cheney's minimalist philosophy, it's a pragmatic middle ground:

- **Better than the old log package** - Structured logging is a significant improvement
- **Simpler than Logrus/Zap/Zerolog** - But more powerful when you need extensibility
- **A good default choice** - For most Go projects in 1.21+, slog is the recommended starting point

### Using slog Idiomatically (Per Cheney's Principles)

If you follow Cheney's advice while using slog:

```go
// DON'T: Log an error at error level AND return it
if err := doSomething(); err != nil {
    slog.Error("failed to do something", "err", err)
    return err  // Already handled, so why log.Error?
}

// DO: Log as info when you handle the error
if err := planA(); err != nil {
    slog.Info("plan A failed, trying plan B", "err", err)
    planB()
}

// DO: Return errors for caller to handle
if err := criticalOperation(); err != nil {
    return fmt.Errorf("critical operation failed: %w", err)
}

// DO: Use log.Error at the boundary (main or HTTP handler)
if err := app.Run(); err != nil {
    slog.Error("application error", "err", err)
    os.Exit(1)
}
```

### Structured Logging as a Solution

Interestingly, slog partially solves one of Cheney's implicit concerns: **By using key-value pairs instead of log levels as the primary differentiator**, structured logging makes the "which level?" question less critical. You can log contextual information and let your log aggregation/filtering tools (like Splunk, Datadog, or ELK) handle level-based filtering.

---

## References

- Original article: https://dave.cheney.net/2015/11/05/lets-talk-about-logging
- Dave Cheney's blog: https://dave.cheney.net
- slog Documentation: https://pkg.go.dev/log/slog (Go 1.21+)
- Related: The package-level logger anti-pattern
- Related: Stack traces and the errors package
