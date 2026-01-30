# Guiding Principles

[← Back to Index](README.md)

Go's design prioritizes three core principles over performance or concurrency alone.

## Clarity

- Code is written for humans first, machines second. It will be read hundreds of times more than written.
- Prioritize maintainability for long-term software engineering.
- Effective communication of ideas is the most important programming skill.
- Ask yourself: "Will the next person understand my intent?"

**Example:**
```go
// Bad: Unclear intent
if x > 0 && y < 100 {
    process(x, y)
}

// Good: Clear intent
const maxValue = 100
if x > 0 && y < maxValue {
    process(x, y)
}

// Better: Self-documenting
func isValidRange(x, y int) bool {
    const maxValue = 100
    return x > 0 && y < maxValue
}

if isValidRange(x, y) {
    process(x, y)
}
```

## Simplicity

- Simplicity is essential for reliability; complexity leads to unreliable software.
- Avoid over-engineering; simple designs are harder to get right but prevent deficiencies.
- Go aims to control complexity in programming.
- When two solutions solve the same problem, choose the simpler one.
- Complexity budget: every feature added should justify its complexity cost.

**Example:**
```go
// Bad: Over-engineered with unnecessary interface
type Processor interface {
    Process() error
}

type DataProcessor struct {
    data []byte
}

func (d *DataProcessor) Process() error {
    // simple processing
    return processData(d.data)
}

// Good: Simple and direct
func processData(data []byte) error {
    // simple processing
    return nil
}
```

## Productivity

- Focus on developer efficiency: fast compilation, readable code, and minimal tooling friction.
- Go enforces consistent formatting (via `gofmt`) to reduce cognitive load and spot errors visually.
- Enables "software engineering at scale" with quick iterations and confidence in changes.

---

**Next:** [Identifiers and Naming →](02_identifiers_naming.md)
