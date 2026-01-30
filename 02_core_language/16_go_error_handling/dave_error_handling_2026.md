# Dave Cheney's Philosophy on Error Handling in Go
## A Technical Deep Dive — Updated for 2026

*Based on the writings of Dave Cheney (2012–2019), updated with modern Go practices through Go 1.22+*

---

## Introduction

Dave Cheney, a prominent figure in the Go community, has written extensively about error handling over more than a decade. His philosophy can be distilled into several core principles that have shaped how the Go community thinks about errors. This document synthesizes his key articles and updates them with the language features introduced since Go 1.13.

---

## Part 1: Why Go Gets Errors Right

### The Historical Context

Dave Cheney's 2012 article *"Why Go Gets Exceptions Right"* sets the stage by examining how different languages have approached error handling:

**C's Approach**: Single return values forced awkward conventions like passing pointers and checking return codes. The infamous `errno` global variable created thread-safety nightmares.

**C++ Exceptions**: Introduced the ability to throw exceptions anywhere, but created the problem that any function could potentially throw without warning. This led to complex patterns like RAII and exception-safe code requirements.

**Java's Checked Exceptions**: Attempted to solve C++'s problem by requiring explicit exception declarations. However, this created backward compatibility nightmares and the dual system of checked/unchecked exceptions proved problematic for language evolution.

### Go's Solution: Errors Are Values

Go takes a fundamentally different approach: errors are simply values returned from functions. This design has several key advantages:

```go
func DoSomething() (Result, error) {
    // ...
}

result, err := DoSomething()
if err != nil {
    return err
}
```

**Key insight from Rob Pike**: Error values in Go aren't special—they're just values like any other. You have the entire language at your disposal to work with them.

**Key insight from Andrei Alexandrescu**: Exceptions are inherently serial. Only one exception can be in flight at a time, requiring immediate and exclusive attention. Go's error values don't have this limitation—you can handle data *and* errors together.

Consider reading from an `io.Reader`:

```go
func ReadAll(r io.Reader) ([]byte, error) {
    var buf = make([]byte, 1024)
    var result []byte
    for {
        n, err := r.Read(buf)
        result = append(result, buf[:n]...)
        if err == io.EOF {
            return result, nil
        }
        if err != nil {
            return nil, err
        }
    }
}
```

This pattern—processing partial data even when an error occurs—is natural in Go but nearly impossible with exception-based systems.

### Panic: Reserved for the Truly Exceptional

Go does have `panic`, but it's fundamentally different from exceptions:

- When you `throw` an exception, you're making it someone else's problem
- When you `panic` in Go, you're declaring game over—the program cannot reasonably continue

```go
// panic is for truly unrecoverable situations
panic("invariant violated: this should never happen")
```

Panics should not be used for error handling. They exist for programming bugs, not operational errors.

---

## Part 2: The Three Strategies for Error Handling

### Strategy 1: Sentinel Errors

Sentinel errors are predeclared error values used for comparison:

```go
if err == io.EOF { /* ... */ }
```

**Problems with sentinel errors:**

1. **Inflexible**: Using `fmt.Errorf` to add context breaks equality checks
2. **API surface expansion**: Public error values become part of your API
3. **Creates coupling**: Packages must import each other to compare errors
4. **Import loops**: Large projects using this pattern risk circular dependencies

**Dave's recommendation**: Avoid sentinel errors. The standard library has some (`io.EOF`), but don't emulate this pattern.

#### Modern Update: Constant Errors

In his 2016 article *"Constant Errors"*, Dave explored an alternative approach:

```go
type Error string

func (e Error) Error() string { return string(e) }

const ErrNotFound = Error("not found")
```

This creates truly constant, immutable error values. Two `Error` values with the same content are equal (fungible), unlike `errors.New` which creates distinct values:

```go
const err1 = Error("EOF")
const err2 = Error("EOF")
fmt.Println(err1 == err2) // true

// Compare to errors.New:
e1 := errors.New("EOF")
e2 := errors.New("EOF")
fmt.Println(e1 == e2) // false
```

### Strategy 2: Error Types

Custom error types provide more context:

```go
type PathError struct {
    Op   string
    Path string
    Err  error
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}
```

Callers can use type assertions:

```go
if pe, ok := err.(*os.PathError); ok {
    fmt.Println("failed during", pe.Op)
}
```

**Problems with error types:**

1. Types must be public for external assertion
2. Creates strong coupling between packages
3. Implementations are constrained to return specific types

**Dave's recommendation**: Avoid error types as part of your public API.

### Strategy 3: Opaque Errors (Preferred)

The most flexible strategy treats errors as opaque—you know an error occurred, but you don't inspect its internals:

```go
func fn() error {
    x, err := bar.Foo()
    if err != nil {
        return err  // Just return it, don't inspect
    }
    // use x
}
```

This creates minimal coupling. The author of `Foo` can change error details without breaking callers.

#### Assert Behavior, Not Type

When you *must* inspect errors, assert for behavior rather than type:

```go
type temporary interface {
    Temporary() bool
}

func IsTemporary(err error) bool {
    te, ok := err.(temporary)
    return ok && te.Temporary()
}
```

This works without importing the package that defines the error or knowing anything about its underlying type—you're simply interested in its *behavior*.

---

## Part 3: Handling Errors Gracefully

### Don't Just Check—Handle

Consider this antipattern:

```go
func AuthenticateRequest(r *Request) error {
    err := authenticate(r.User)
    if err != nil {
        return err
    }
    return nil
}
```

If `authenticate` fails deep in the call stack, the top-level error might just be `"no such file or directory"` with no indication of *where* or *why*.

### Adding Context with Wrapping

Dave championed the `github.com/pkg/errors` package, which introduced wrapping:

```go
func ReadFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, errors.Wrap(err, "open failed")
    }
    defer f.Close()

    buf, err := ioutil.ReadAll(f)
    if err != nil {
        return nil, errors.Wrap(err, "read failed")
    }
    return buf, nil
}
```

Output: `could not read config: open failed: open /path/to/file: no such file or directory`

#### Modern Update: Go 1.13+ Error Wrapping

The standard library now supports error wrapping natively with `%w`:

```go
func ReadFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("open failed: %w", err)
    }
    defer f.Close()

    buf, err := io.ReadAll(f)
    if err != nil {
        return nil, fmt.Errorf("read failed: %w", err)
    }
    return buf, nil
}
```

**Inspecting wrapped errors:**

```go
// Check if an error matches a specific value
if errors.Is(err, os.ErrNotExist) {
    // handle missing file
}

// Extract a specific error type
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("operation:", pathErr.Op)
}
```

**Important**: Use `%w` when you want to expose the underlying error. Use `%v` when the underlying error is an implementation detail. Wrapping makes the error part of your API contract.

### Only Handle Errors Once

Handling an error means inspecting it and making a decision. Making *less* than one decision ignores the error. Making *more* than one decision creates problems:

```go
// BAD: handles the error twice
func Write(w io.Writer, buf []byte) error {
    _, err := w.Write(buf)
    if err != nil {
        log.Println("unable to write:", err)  // Decision 1: log
        return err                             // Decision 2: return
    }
    return nil
}
```

This results in duplicate log entries with the original error stripped of context at the top of the stack.

**Better approach:**

```go
func Write(w io.Writer, buf []byte) error {
    _, err := w.Write(buf)
    return errors.Wrap(err, "write failed")
}
```

Wrap the error with context and let the caller decide how to handle it.

---

## Part 4: Eliminating Errors

### The Best Error Handling Is No Error Handling

Dave's 2019 article *"Eliminate Error Handling by Eliminating Errors"* draws from John Ousterhout's *A Philosophy of Software Design*.

**Before: Error handling obscures intent**

```go
func CountLines(r io.Reader) (int, error) {
    var (
        br    = bufio.NewReader(r)
        lines int
        err   error
    )
    for {
        _, err = br.ReadString('\n')
        lines++
        if err != nil {
            break
        }
    }
    if err != io.EOF {
        return 0, err
    }
    return lines, nil
}
```

**After: Clean API eliminates error handling**

```go
func CountLines(r io.Reader) (int, error) {
    sc := bufio.NewScanner(r)
    lines := 0
    for sc.Scan() {
        lines++
    }
    return lines, sc.Err()
}
```

`bufio.Scanner` encapsulates the complexity of handling `io.EOF` vs. real errors, leaving clean, readable code.

### The errWriter Pattern

For repetitive I/O operations, create helper types that accumulate errors:

```go
type errWriter struct {
    io.Writer
    err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
    if e.err != nil {
        return 0, e.err
    }
    var n int
    n, e.err = e.Writer.Write(buf)
    return n, nil
}
```

**Usage:**

```go
func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
    ew := &errWriter{Writer: w}
    fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
    for _, h := range headers {
        fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
    }
    fmt.Fprint(ew, "\r\n")
    io.Copy(ew, body)
    return ew.err
}
```

No error checking after each operation—just check once at the end.

---

## Part 5: Modern Go Error Features (2019–2026)

### Go 1.13: Error Wrapping in the Standard Library

The `%w` verb in `fmt.Errorf` and the `errors.Is`/`errors.As` functions standardized patterns from `github.com/pkg/errors`:

```go
// Wrap with context
err := fmt.Errorf("user %d: %w", userID, ErrNotFound)

// Check error chain
if errors.Is(err, ErrNotFound) {
    // handle not found
}

// Extract typed error
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    fmt.Println("field:", validationErr.Field)
}
```

### Go 1.20: Multiple Error Wrapping

Go 1.20 introduced `errors.Join` and support for multiple `%w` verbs:

```go
// Join multiple errors
err := errors.Join(err1, err2, err3)

// Multiple wrapping in fmt.Errorf
err := fmt.Errorf("operation failed: %w, %w", errA, errB)
```

**Traversing joined errors:**

```go
err := errors.Join(ErrPermission, ErrTimeout)

// errors.Is checks the entire tree
fmt.Println(errors.Is(err, ErrPermission)) // true
fmt.Println(errors.Is(err, ErrTimeout))    // true

// Access the slice directly via type assertion
if je, ok := err.(interface{ Unwrap() []error }); ok {
    for _, e := range je.Unwrap() {
        fmt.Println(e)
    }
}
```

**Use cases:**

- Collecting errors from multiple goroutines
- Cleanup errors combined with operation errors
- Validation errors from multiple fields

### Custom Error Types with Unwrap

Implement `Unwrap` on your custom errors to participate in the chain:

```go
type DatabaseError struct {
    Op  string
    Err error
}

func (d *DatabaseError) Error() string {
    return fmt.Sprintf("database error during %s: %v", d.Op, d.Err)
}

func (d *DatabaseError) Unwrap() error {
    return d.Err
}
```

For joined errors, implement `Unwrap() []error` instead.

---

## Summary: Dave Cheney's Error Handling Principles

1. **Errors are values**—treat them as such, using the full power of the language

2. **Avoid sentinel errors**—they create coupling and expand your API surface

3. **Avoid public error types**—they constrain implementations and create brittle APIs

4. **Prefer opaque errors**—return errors without assuming anything about their contents

5. **Assert behavior, not type**—use interfaces like `Temporary()` to check capabilities

6. **Add context when wrapping**—use `fmt.Errorf("context: %w", err)` to preserve the chain

7. **Handle errors only once**—either log or return, never both

8. **Eliminate errors through design**—APIs like `bufio.Scanner` can remove error handling burden

9. **Use `errors.Is` and `errors.As`**—modern Go provides robust inspection without type coupling

10. **Wrap judiciously**—wrapping exposes errors as part of your API; use `%v` for implementation details

---

## References

- [Why Go Gets Exceptions Right](https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right) (2012)
- [Error Handling vs. Exceptions Redux](https://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux) (2014)
- [Inspecting Errors](https://dave.cheney.net/2014/12/24/inspecting-errors) (2014)
- [Errors and Exceptions, Redux](https://dave.cheney.net/2015/01/26/errors-and-exceptions-redux) (2015)
- [Constant Errors](https://dave.cheney.net/2016/04/07/constant-errors) (2016)
- [Don't Just Check Errors, Handle Them Gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) (2016)
- [Eliminate Error Handling by Eliminating Errors](https://dave.cheney.net/2019/01/27/eliminate-error-handling-by-eliminating-errors) (2019)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) — Official Go Blog
- [Go 1.20 Release Notes: errors.Join](https://go.dev/doc/go1.20)

---

## Part 6: Error Handling in 2026 — The State of the Art

### The Go Team's Decision: No New Syntax

In June 2025, Robert Griesemer published a landmark blog post titled *"[ On | No ] syntactic support for error handling"* that officially closed the door on new error handling syntax for the foreseeable future.

After years of proposals—the 2018 `check`/`handle` draft, the 2019 `try` proposal (which garnered nearly 900 comments), and the 2024 `?` operator proposal—the Go team concluded that no approach had achieved sufficient consensus. The decision was pragmatic: without broad agreement, language changes cause more harm than good.

**Key arguments from the Go team:**

1. **The opportunity has passed**: Go is 15+ years old with a mature ecosystem. Introducing syntactic sugar now would fracture the community between those who want change and those who prefer the status quo.

2. **Tooling compensates**: Modern IDEs with AI-assisted code completion make writing error checks trivial. IDEs could even provide toggle switches to collapse error handling code during reading.

3. **Verbosity signals handling**: The `if err != nil` pattern, while verbose, makes error handling highly visible to code reviewers. Hidden control flow (like `try` or `?`) can obscure what happens when things go wrong.

4. **Debugging benefits**: Having explicit `if` statements means you can easily add `println` calls or set debugger breakpoints without restructuring code.

5. **Real handling isn't verbose**: When errors are properly handled (not just returned), the boilerplate becomes a smaller fraction of the code.

### Dave Cheney's Philosophy Vindicated

The 2025 decision aligns remarkably well with Dave Cheney's long-standing positions:

- **Errors are values**: The Go team explicitly referenced Rob Pike's 2015 blog post and the power of treating errors as regular values
- **Design errors out of existence**: Rather than syntactic changes, the team pointed to library-level solutions
- **Opaque handling works**: The lack of special syntax forces developers to treat errors explicitly

### What's New in Go 1.22–1.24 for Errors

While syntax hasn't changed, the ecosystem has evolved:

**Go 1.22: `cmp.Or` for Error Coalescing**

```go
func printSum(a, b string) error {
    x, err1 := strconv.Atoi(a)
    y, err2 := strconv.Atoi(b)
    if err := cmp.Or(err1, err2); err != nil {
        return err
    }
    fmt.Println("result:", x+y)
    return nil
}
```

**Go 1.23: Iterators and Error Handling Patterns**

The `iter` package introduced standardized iterator patterns. For error-prone iteration, the community has adopted the `bufio.Scanner` pattern with a separate `Err()` method:

```go
type Playlist struct {
    list    []Song
    indices []int
    err     error
}

func (p *Playlist) All(yield func(int, Song) bool) {
    for _, index := range p.indices {
        if index < 0 || index >= len(p.list) {
            p.err = fmt.Errorf("index %d out of bounds", index)
            return
        }
        if !yield(index, p.list[index]) {
            return
        }
    }
}

func (p *Playlist) Err() error { return p.err }

// Usage:
for idx, song := range playlist.All {
    fmt.Printf("Playing: %s\n", song.Title)
}
if err := playlist.Err(); err != nil {
    log.Fatal(err)
}
```

This is exactly the pattern Dave Cheney praised in `bufio.Scanner`—eliminating error handling by encapsulating it in the API design.

**Go 1.23: `iter.Seq2` for Error-Yielding Iterators**

For cases where errors must be yielded inline:

```go
func ParseNumbers(strs []string) iter.Seq2[int, error] {
    return func(yield func(int, error) bool) {
        for _, s := range strs {
            n, err := strconv.Atoi(s)
            if !yield(n, err) {
                return
            }
        }
    }
}

// Usage:
for num, err := range ParseNumbers(input) {
    if err != nil {
        break  // Stop at first error
    }
    process(num)
}
```

### 2025 Developer Survey Insights

The January 2026 Go Developer Survey results revealed interesting findings:

- **91% satisfaction** with Go overall
- **Error handling remains the top complaint**, but enthusiasm for syntax changes has waned
- Developers new to Go from exception-based languages feel the pain most acutely
- Experienced Go developers often report the issue becomes less important as they write more idiomatic code
- The survey noted opportunities for **context-specific guidance** (e.g., "Error handling in Go for Java developers")

### Best Practices for 2026

Based on Dave Cheney's principles, modern Go features, and community consensus:

**1. Wrap errors with context using `%w`**
```go
if err != nil {
    return fmt.Errorf("loading config from %s: %w", path, err)
}
```

**2. Use `errors.Is` and `errors.As` for inspection**
```go
if errors.Is(err, os.ErrNotExist) {
    // Handle missing file
}

var pathErr *os.PathError
if errors.As(err, &pathErr) {
    log.Printf("operation %s failed on %s", pathErr.Op, pathErr.Path)
}
```

**3. Use `errors.Join` for multiple concurrent errors**
```go
var errs []error
var wg sync.WaitGroup
var mu sync.Mutex

for _, url := range urls {
    wg.Add(1)
    go func(u string) {
        defer wg.Done()
        if err := fetch(u); err != nil {
            mu.Lock()
            errs = append(errs, err)
            mu.Unlock()
        }
    }(url)
}
wg.Wait()

if err := errors.Join(errs...); err != nil {
    return err
}
```

**4. Design APIs that eliminate error handling**

Following Dave Cheney's advice, use patterns like `bufio.Scanner`:

```go
// Instead of:
for {
    line, err := reader.ReadLine()
    if err == io.EOF {
        break
    }
    if err != nil {
        return err
    }
    process(line)
}

// Prefer:
scanner := bufio.NewScanner(reader)
for scanner.Scan() {
    process(scanner.Text())
}
return scanner.Err()
```

**5. Use helper types for repetitive I/O**

```go
type errWriter struct {
    w   io.Writer
    err error
}

func (ew *errWriter) Write(p []byte) {
    if ew.err != nil {
        return
    }
    _, ew.err = ew.w.Write(p)
}

func (ew *errWriter) Printf(format string, args ...any) {
    if ew.err != nil {
        return
    }
    _, ew.err = fmt.Fprintf(ew.w, format, args...)
}
```

**6. Assert behavior, not type**

```go
type temporary interface {
    Temporary() bool
}

func shouldRetry(err error) bool {
    var t temporary
    return errors.As(err, &t) && t.Temporary()
}
```

### The Future

The Go team has stated they will close all incoming proposals focused primarily on error handling syntax without further investigation. This doesn't mean error handling is a solved problem—it means the community should focus energy elsewhere:

- **Better tooling**: IDE features to collapse/expand error handling blocks
- **Better libraries**: APIs designed to minimize error handling burden
- **Better documentation**: Context-specific guidance for developers from different backgrounds
- **Better static analysis**: Linters that catch improper error handling

Dave Cheney's philosophy has become the official position: errors are values, treat them with care, and design your APIs to make handling them natural rather than burdensome.

---

## Conclusion

Dave Cheney's decade of writing on Go error handling has proven prescient. His core insights—that errors are values, that opaque handling provides flexibility, that API design can eliminate error handling burden—have become the foundation of modern Go best practices.

The Go team's 2025 decision to abandon syntactic changes validates his approach: rather than adding special syntax, Go trusts developers to use the full power of the language to handle errors appropriately. The additions since Go 1.13 (`%w` wrapping, `errors.Is`/`errors.As`, `errors.Join`) provide the tools needed without special syntax.

For Go developers in 2026, the path forward is clear:

1. Embrace explicit error handling as a feature, not a bug
2. Use wrapping to preserve context through the call stack
3. Design APIs that encapsulate error handling complexity
4. Assert behavior rather than types for maximum flexibility
5. Handle errors once, at the appropriate level

As Dave Cheney wrote: *"Errors are part of your package's public API—treat them with as much care as you would any other part of your public API."*

---

## References

- [Why Go Gets Exceptions Right](https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right) (2012)
- [Error Handling vs. Exceptions Redux](https://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux) (2014)
- [Inspecting Errors](https://dave.cheney.net/2014/12/24/inspecting-errors) (2014)
- [Errors and Exceptions, Redux](https://dave.cheney.net/2015/01/26/errors-and-exceptions-redux) (2015)
- [Constant Errors](https://dave.cheney.net/2016/04/07/constant-errors) (2016)
- [Don't Just Check Errors, Handle Them Gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) (2016)
- [Eliminate Error Handling by Eliminating Errors](https://dave.cheney.net/2019/01/27/eliminate-error-handling-by-eliminating-errors) (2019)
- [Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors) — Official Go Blog
- [Go 1.20 Release Notes: errors.Join](https://go.dev/doc/go1.20)
- [Go 1.23 Release Notes: Iterators](https://go.dev/blog/go1.23)
- [[ On | No ] syntactic support for error handling](https://go.dev/blog/error-syntax) — Robert Griesemer, June 2025
- [2025 Go Developer Survey Results](https://go.dev/blog/survey2025) — January 2026

---

*Last updated: January 2026*
