# Identifiers and Naming

[← Back to Index](README.md) | [Next: Comments and Documentation →](03_comments_documentation.md)

Good naming is crucial due to Go's minimal syntax; names must convey intent clearly.

## Choose Identifiers for Clarity, Not Brevity

- Optimize for readability, not line count or typing speed.
- Qualities of good names: concise, descriptive, predictable, and idiomatic.
- Describe purpose/behavior/result, not implementation or contents.

## Identifier Length

- Use the "right" length: short for local variables (e.g., `p` in loops), longer for broader scope.
- Avoid type names in variables (e.g., no `personSlice` for `[]Person`).
- Constants describe values, not usage.
- Guidelines: single letters for loops/branches, single words for params/returns, multiple for functions/packages.

**Examples by Scope:**
```go
// Loop indices: single letters
for i := 0; i < len(items); i++ {
    // ...
}

// Short-lived locals: abbreviated but clear
for _, p := range products {
    processProduct(p)
}

// Function parameters: single word
func calculateTotal(price, quantity int) int {
    return price * quantity
}

// Package-level or long-lived: descriptive
var defaultConnectionTimeout = 30 * time.Second

// Bad: Type in name
var userSlice []User      // Don't
var userMap map[int]User  // Don't

// Good: Describe the collection
var users []User
var usersByID map[int]User
```

## Consistent Naming and Declaration Style

- Follow Go conventions (e.g., camelCase, exported names capitalized).
- Use consistent declarations (e.g., `var` vs. `:=`).
- Be a team player: match project style for collaboration.

---

[← Previous: Guiding Principles](01_guiding_principles.md) | [Next: Comments and Documentation →](03_comments_documentation.md)
