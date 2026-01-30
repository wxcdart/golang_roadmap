# How to Name Things: The Hardest Problem in Programming

*Based on Peter Hilton's presentation, adapted for Go conventions and updated for 2026*

---

## The Core Problem

> "There are only two hard things in Computer Science: cache invalidation and naming things."
> — Phil Karlton

Naming matters because code is read far more often than it's written. Good names reduce cognitive load; bad names create confusion and bugs.

---

## Orwell's Rules, Adapted for Go

George Orwell's 1946 essay "Politics and the English Language" offers timeless writing advice. Here's how it applies to Go:

### 1. Avoid Clichéd Patterns

**Orwell**: Never use a figure of speech you're used to seeing in print.

**For Go**: Don't reflexively add pattern suffixes. Go's standard library avoids `AbstractConfigurationFactory` bloat.

```go
// Avoid
type AbstractServiceFactory struct{}

// Prefer
type Dialer struct{}  // what it does, not what pattern it is
```

### 2. Prefer Short Words — Go Takes This Seriously

**Orwell**: Never use a long word where a short one will do.

**For Go**: This is idiomatic. Short names are not just acceptable but *preferred* for local scope.

```go
// Idiomatic Go
for i, v := range items {}
func Copy(dst Writer, src Reader) {}

// Not idiomatic
for index, value := range items {}
func Copy(destination Writer, source Reader) {}
```

**The Go proverb**: "The greater the distance between a name's declaration and its uses, the longer the name should be."

### 3. Cut Unnecessary Words

**Orwell**: If it is possible to cut a word out, always cut it out.

**For Go**: Absolutely. Avoid redundancy with package names.

```go
// Avoid (stuttering)
http.HTTPServer
context.ContextKey

// Prefer
http.Server
context.Key
```

### 4. Use Active Voice

**Orwell**: Never use the passive where you can use the active.

**For Go**: Name interfaces by what they *do*, using -er suffix.

```go
// Good: describes the action
type Reader interface { Read(p []byte) (n int, err error) }
type Stringer interface { String() string }

// Avoid: passive or noun-heavy
type DataContainer struct{}
```

### 5. Avoid Unnecessary Jargon

**Orwell**: Never use jargon if you can think of an everyday equivalent.

**For Go**: Keep domain terms; drop computer science theater.

```go
// Avoid
type ShipmentMonad struct{}
type UserEntityBean struct{}

// Prefer
type Shipment struct{}
type User struct{}
```

### 6. Break Rules to Avoid Barbarism

**Orwell**: Break any of these rules sooner than say anything outright barbarous.

**For Go**: Follow conventions, but prioritize clarity. If a longer name genuinely helps, use it.

---

## Go-Specific Naming Conventions

### Variable Name Length by Scope

| Scope | Style | Example |
|-------|-------|---------|
| Loop index | Single letter | `i`, `j`, `k` |
| Local (tight scope) | Short | `r`, `w`, `buf`, `ctx` |
| Local (wider scope) | Descriptive | `userCount`, `reqBody` |
| Package-level | Descriptive | `DefaultTimeout` |
| Exported | Clear & complete | `ResponseWriter` |

### Standard Abbreviations (Use These)

```go
ctx  // context.Context
req  // request
resp // response
err  // error
buf  // buffer
cfg  // config
srv  // server
conn // connection
msg  // message
arg  // argument
```

### Acronyms: ALL CAPS

```go
// Correct
var userID string
type HTTPClient struct{}
func ParseXML() {}

// Incorrect
var oderId string
type HttpClient struct{}
```

### Receiver Names: Short and Consistent

```go
// Good: short, typically first letter(s) of type
func (s *Server) ListenAndServe() error {}
func (c *Client) Do(req *Request) (*Response, error) {}

// Avoid: generic or long
func (this *Server) ListenAndServe() error {}
func (server *Server) ListenAndServe() error {}
```

### Interface Naming

```go
// Single-method: use -er suffix
type Reader interface { Read([]byte) (int, error) }
type Closer interface { Close() error }

// Multi-method: describe the capability
type ReadCloser interface {
    Reader
    Closer
}
```

---

## Universal Naming Sins (Still Apply in Go)

### The Worst Names

| Bad Name | Why | Better |
|----------|-----|--------|
| `data` | Meaningless | `payload`, `record` |
| `data2` | Even worse | Rethink your design |
| `info` | Too vague | `metadata`, `stats` |
| `tmp` | Lazy | Name what it holds |
| `Manager` | Vague catch-all | `Pool`, `Registry`, `Scheduler` |
| `utils` | Junk drawer | Split by actual purpose |

### Replace Vague Verbs

| Vague | More Precise Alternatives |
|-------|---------------------------|
| `get` | `fetch`, `find`, `lookup`, `load`, `read` |
| `do` | `execute`, `process`, `run`, `apply` |
| `handle` | `route`, `dispatch`, `process` |
| `manage` | `schedule`, `coordinate`, `pool` |

Use `get` only for simple field accessors. If it does I/O, use `fetch` or `load`.

---

## Naming in 2026: Context-Aware Naming and Tooling

### AI-Assisted Code Review

Modern IDEs and GitHub Copilot can suggest better names during code review:

```go
// IDE/AI notices the vague name and suggests alternatives
users := GetData()     // Suggestion: GetUsers(), FetchUsers()
result := ProcessIt()  // Suggestion: Parse(), Transform(), Validate()
```

### Automated Naming Checks

2026 tooling has improved significantly:

1. **`golangci-lint`** with naming rules catches obvious issues
2. **gopls** provides real-time rename suggestions
3. **Code review bots** can flag naming improvements before merge
4. **Custom analyzers** can enforce domain-specific naming conventions

```bash
# Modern tooling catches weak names
golangci-lint run ./...  # includes naming checks

# IDE rename with full-codebase awareness
gopls rename  # integrated in VS Code, GoLand
```

### Domain-Driven Naming

As teams mature, establish a *ubiquitous language*:

```go
// E-commerce domain example
type Order struct{}           // Not "Transaction"
type ShoppingCart struct{}    // Not "Basket" or "Trolley"
type Fulfillment struct{}     // Not "Shipping"
type Reconciliation struct{}  // Not "SettlementProcessing"

// This vocabulary comes from business stakeholders, not programmers
```

---

## Practical Advice

### How to Get Better at Naming

1. **Read good Go code** — stdlib, well-maintained OSS projects
2. **Use `go vet` and linters** — they catch naming issues
3. **Review diffs for naming** — make it part of code review
4. **Rename fearlessly** — `gorename` and IDE tools make it safe
5. **Build domain vocabulary** — learn what your users call things

### The Rename Refactor

Renaming is the simplest but most effective refactoring. In Go:

```bash
# Using gorename
gorename -from '"mypackage".oldName' -to newName

# Or use your IDE's rename function (F2 in VS Code)
```

### When in Doubt

Ask yourself:
1. Would a new team member understand this name?
2. Does it match Go conventions?
3. Is it as short as possible while remaining clear?
4. Does it avoid stuttering with the package name?

---

## Summary

| Hilton's Advice | Go Alignment |
|-----------------|--------------|
| Avoid clichéd patterns | ✅ Go avoids pattern-heavy naming |
| Prefer short words | ✅✅ Go takes this further than most |
| Cut unnecessary words | ✅ Don't stutter with package names |
| Use active voice | ✅ Interface -er suffix convention |
| Avoid jargon | ✅ Keep it simple |
| Know when to break rules | ✅ Clarity wins |

**The Go philosophy**: Local variables should be short. Exported names should be clear. The distance between declaration and use determines name length. When in doubt, shorter is more idiomatic.

---

*"Clear is better than clever."* — Go Proverb

---

## References

### Original Presentation

- Peter Hilton — How to Name Things: The Hardest Problem in Programming (2014)
- George Orwell — Politics and the English Language (1946)

### Go Naming Conventions

- Effective Go — Names — Official Go documentation on naming
- Go Code Review Comments — Naming — Community conventions from the Go wiki
- Go Blog — Package Names — Andrew Gerrand on naming packages
- Go Proverbs — go-proverbs.github.io — Rob Pike's guiding principles

### Books Referenced by Hilton

- Steve McConnell — Code Complete, Chapter 10: The Power of Variable Names
- Robert C. Martin — Clean Code, Chapter 2: Meaningful Names
- Eric Evans — Domain-Driven Design, Chapter 2: Communication and the Use of Language

### Tools

- Codelf — github.com/unbug/codelf — Search real-world variable names from open source projects
- gorename — golang.org/x/tools/cmd/gorename — Safe renaming tool for Go
- gopls — golang.org/x/tools/gopls — Go language server with rename support
- golangci-lint — golangci-lint.run — Comprehensive linting with naming rules

---

*Last updated: January 2026*
