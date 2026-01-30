# Go Has Both make and new Functions, What Gives?

Based on the article by Dave Cheney: https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives

---

## The Apparent Redundancy

Go has many ways to initialize variables, which can create confusion. Rob Pike noted at Gophercon that several ways appear to do the same thing:

```go
s := &SomeStruct{}
v := SomeStruct{}
s := &v              // identical to the first line
s := new(SomeStruct) // also identical
```

It's fair to ask: why have both `make` and `new`? They seem redundant. But they actually serve different purposes.

---

## The Key Difference

**`make`** and **`new`** do fundamentally different things:

- **`new(T)`** allocates memory for any type `T` and returns a pointer to it (`*T`). The value is initialized to the zero value.
- **`make(T)`** only works with slices, maps, and channels. It allocates and initializes these types **and returns the value itself, not a pointer**.

---

## Why Can't We Use make for Everything?

### The Problem with Generic Types

Go doesn't have user-defined generic types, but it has three built-in "generic" types that operate as collections:
- **Slices** - ordered lists
- **Maps** - key-value stores
- **Channels** - communication primitives

**`make` is specifically designed for these three types.** It must be provided by the runtime because there's no way to express `make`'s function signature directly in Go—it has special behavior that varies depending on the type.

### Why `make` Doesn't Return Pointers

Although `make` creates values of generic types, they are still just regular values. **`make` does not return pointer values.**

If `new` was removed in favor of `make`, how would you construct a pointer to an initialized value?

```go
var x1 *int
var x2 = new(int)
```

Both `x1` and `x2` have the same type: `*int`. But:
- `x1` is `nil` and unsafe to dereference
- `x2` points to initialized memory (the zero value `0`) and is safe to dereference

---

## Why Can't We Use new for Everything?

Although `new` is rarely used, its behavior is well-specified:

**`new(T)` always returns `*T` pointing to an initialized `T`.** Since Go doesn't have constructors, the value is initialized to `T`'s zero value.

Using `new` to construct pointers to slices, maps, or channels works and is consistent with `new`'s semantics:

```go
// Slice
s := new([]string)
fmt.Println(len(*s))  // 0
fmt.Println(*s == nil) // true

// Map
m := new(map[string]int)
fmt.Println(m == nil)  // false
fmt.Println(*m == nil) // true

// Channel
c := new(chan int)
fmt.Println(c == nil)  // false
fmt.Println(*c == nil) // true
```

---

## Could We Merge Them? (No, Here's Why)

**Could `new` be extended to operate like `make` for slices, maps, and channels?**

Technically yes, but it would introduce its own inconsistencies and complications:

### 1. Special Behavior Rule
`new` would have special behavior only when the type is a slice, map, or channel. This is a special case every Go programmer would have to remember and apply differently depending on context.

### 2. Variable Arguments Problem
For slices and channels, `new` would need to become variadic to accept optional parameters:
- Slice length/capacity
- Channel buffer size
- Map initial capacity

Currently, `new` takes exactly one argument: the type. Adding variadic arguments would introduce more special cases to learn and remember.

### 3. Return Type Confusion
`new` always returns `*T` for the `T` passed to it. If `new` returned a slice value (not a pointer), code like this would break:

```go
func Read(buf []byte) []byte
// assume new takes an optional length
buf := Read(new([]byte, 4096))  // Error: new would return []byte, not *[]byte
```

You'd need `*new([]byte, length)` or other special syntax to make it work—yet more special cases in the grammar.

---

## Summary

`make` and `new` do different things:

- **`new(T)`** - Allocates memory for any type, returns a pointer to initialized memory (`*T`)
- **`make(T)`** - Only for slices, maps, and channels; returns the initialized value (not a pointer)

### Advice for Go Programmers

1. **Use `new` sparingly** - There are almost always easier or cleaner ways to write your program without it. Prefer struct literals: `&SomeStruct{}` is clearer than `new(SomeStruct)`.

2. **Use `make` for collections** - It's the idiomatic way to create slices, maps, and channels.

3. **Code review signal** - Use of `new`, like use of named return arguments, is a signal that code is trying to do something clever. Pay attention! It may be clever and justified, but more likely it can be rewritten more clearly and idiomatically.

### Coming from Other Languages

If you're coming from a language that uses constructors (like Java or C++), it may seem like `new` should be all you need. But Go is different: it doesn't have constructors, and its initialization philosophy is based on zero values and simplicity.

---

## References

- Original article: https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives
- Related: What is the zero value, and why is it useful?
- Related: On declaring variables
