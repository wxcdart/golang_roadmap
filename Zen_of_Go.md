# The Zen of Go

Based on the article by Dave Cheney: https://dave.cheney.net/2020/02/23/the-zen-of-go

## Introduction

Go, like many programming languages, has developed a set of engineering values that guide how Gophers write code. This document explores those values by drawing parallels with Tim Peters' "The Zen of Python" (PEP-20) and adapting them to Go's philosophy.

---

## 1. A Good Package Starts with a Good Name

> "Namespaces are one honking great idea–let's do more of those!" - The Zen of Python, Item 19

In Go, packages are namespaces. Each package should have a single, clear purpose described by its name—a noun that conveys what the package provides.

**Why?** Because change is inevitable in software. Design is the art of arranging code to work today and be changeable forever. When a package's name no longer matches its responsibility, it's time to replace it rather than force it to do more.

> "Design is the art of arranging code to work today, and be changeable forever." - Sandi Metz

A good package name is like an elevator pitch in a single word. It tells you exactly what the package provides.

---

## 2. Simplicity Matters

> "Simple is better than complex." - The Zen of Python, Item 3

Go holds simplicity as a core value. While simplicity is not the same as ease, simple code is:
- **Readable** and maintainable
- **Reliable** and understandable
- Not crude or unsophisticated

> "There are two ways of constructing a software design: One way is to make it so simple that there are obviously no deficiencies, and the other way is to make it so complicated that there are no obvious deficiencies. The first method is far more difficult." - C. A. R. Hoare

> "Simplicity is prerequisite for reliability." - Edsger W Dijkstra

> "Controlling complexity is the essence of computer programming." - Brian W. Kernighan

When it comes to Go, simple code is preferable to clever code.

---

## 3. Avoid Package Level State

> "Explicit is better than implicit." - The Zen of Python, Item 2

Being explicit in Go means being clear about **coupling** and **state**. 

**Coupling** measures how much one thing depends on another. Tightly coupled code is harder to test, harder to parallelize, and harder to reuse.

**State** is the core problem in computer science. Package-level state creates hidden dependencies and makes code fragile.

### Example Problem:
```go
package counter

var count int

func Increment(n int) int {
    count += n
    return count
}
```

This counter cannot be tested in isolation, tested in parallel, or used multiple times per program.

### Solution:
```go
package counter

type Counter struct {
    count int
}

func (c *Counter) Increment(n int) int {
    c.count += n
    return c.count
}
```

**Principle:** Reduce coupling and avoid "spooky action at a distance" by providing dependencies as fields on types rather than using package variables.

---

## 4. Plan for Failure, Not Success

> "Errors should never pass silently." - The Zen of Python, Item 10

Exception-based languages follow the Samurai principle: "return victorious or not at all." Go takes a different approach.

**Go's philosophy:** Robust programs handle failure cases **before** they handle the happy path.

In server programs, multi-threaded programs, and network programs, failures are inevitable:
- Unexpected data
- Timeouts
- Connection failures
- Corrupted data

These must be front and center when writing Go code.

> "I think that error handling should be explicit, this should be a core value of the language." - Peter Bourgon

The verbosity of `if err != nil { return err }` is justified because it forces deliberate thinking about each failure condition at the point where it occurs.

---

## 5. Return Early Rather Than Nesting Deeply

> "Flat is better than nested." - The Zen of Python, Item 5

**"Line of sight" coding** means avoiding control flow that requires deep indentation.

Techniques:
- Use guard clauses to return early if preconditions aren't met
- Place successful returns at the end of functions
- Extract functions and methods to reduce indentation levels

> "Line of sight is a straight line along which an observer has unobstructed vision." - Mat Ryer

Every level of indentation adds another precondition to the programmer's cognitive stack. Keep the successful path of the function close to the left edge of your screen. This prevents your code from sliding out of sight and improves readability.

---

## 6. If You Think It's Slow, Prove It with a Benchmark

> "In the face of ambiguity, refuse the temptation to guess." - The Zen of Python, Item 12

Don't guess about performance. Common dogma claims:
- "defer is slow"
- "CGO is expensive"
- "Always use atomics, not mutexes"

These may be outdated or context-dependent.

**Better approach:** Use Go's built-in benchmarking and profiling tools to find actual bottlenecks. Make data-driven optimization decisions, not faith-based ones.

> "APIs should be easy to use and hard to misuse." - Josh Bloch

---

## 7. Before You Launch a Goroutine, Know When It Will Stop

Goroutines are Go's signature feature for lightweight concurrency. They're so easy to start—just prefix with `go`—but this power comes with responsibility.

**Critical questions before launching a goroutine:**

1. **Under what condition will a goroutine stop?**
   - Go has no way to kill a goroutine; you must ask it politely
   - This is typically done via channel operations
   - Closed channels signal completion

2. **What is required for that condition to arise?**
   - If channels signal completion, who closes the channel and when?

3. **What signal will you use to know the goroutine has stopped?**
   - Use channels to signal completion
   - Use `sync.WaitGroup` for fan-in patterns

> "You type 'go', a space, and then a function call. Three keystrokes, you can't make it much shorter than that. Three keystrokes and you've just started a subprocess." - Rob Pike

Goroutines are cheap but not free. At 10^6 goroutines, memory and resource management become important. Resource ownership is tied to goroutine lifetime—when the goroutine exits, its resources are freed.

---

## 8. Leave Concurrency to the Caller

Library writers should leave the responsibility of starting goroutines to the caller. Let users decide:
- **How** to start functions asynchronously
- **How** to track execution
- **How** to wait for completion

This gives callers flexibility and maintains explicit control over concurrency.

> Example frameworks: go-kit's `run.Group` and Go's `golang.org/x/sync/errgroup`

The principle: **Don't spawn goroutines in libraries; let the caller decide.**

---

## 9. Write Tests to Lock In Your Package's API Behavior

Your tests are the contract of what your software does and does not do.

**Unit tests should:**
- Lock in the behavior of your package's API
- Describe in code what the package promises to do
- Define the contract in code, not documentation

```bash
go test
```

This command lets you verify with high confidence that behavior people relied on before your change continues to work after your change.

**Key principle:** Any change that adds, modifies, or removes a public API must include changes to its tests.

---

## 10. Moderation Is a Virtue

Go has only 25 keywords. Its features stand out precisely because there are so few. Go programmers often become excited and overuse language features:

- **Channels everywhere?** Can lead to hard-to-test, fragile code
- **Too many goroutines?** Creates management complexity
- **Excessive embedding?** Can recreate the fragile base class problem
- **Interfaces for everything?** Over-engineering for flexibility

**Principle:** Use moderation. Prefer simpler approaches over clever ones.

> All things in moderation—Go's features are no exception.

---

## 11. Maintainability Counts

> "Readability Counts." - The Zen of Python, Item 7

Readability is a stepping stone to the real goal: **maintainability**.

Go is not optimized for:
- Clever one-liners
- Minimum line count
- Shortest source code
- Fastest typing

Go optimizes for:
- **Clarity to the reader**
- **Code that can be maintained after the original author is gone**
- **Code that serves as a foundation for future value**

> "Can the thing you worked hard to build be maintained after you're gone? What can you do today to make it easier for someone to maintain your code tomorrow?"

If software cannot be maintained, it will be rewritten. Make your code maintainable so that Go remains a solid investment for your organization and community.

---

## Summary

The Zen of Go emphasizes:
1. Clear package names and single purposes
2. Simplicity over cleverness
3. Explicit over implicit (especially state and coupling)
4. Planning for failure as a first-class concern
5. Readable code structure (line of sight)
6. Data-driven optimization over dogma
7. Responsible goroutine management
8. Caller-controlled concurrency
9. Tests as API contracts
10. Moderation in feature use
11. Maintainability as the ultimate goal

---

## References

- Original article: https://dave.cheney.net/2020/02/23/the-zen-of-go
- The Zen of Python (PEP-20): https://www.python.org/dev/peps/pep-0020/
- Dave Cheney's blog: https://dave.cheney.net
