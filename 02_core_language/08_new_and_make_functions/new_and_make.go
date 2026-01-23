// new_and_make.go
// Demonstrates the difference between new and make in Go.
// Includes idiomatic usage and comments per Effective Go.

package main

import (
	"fmt"
)

func main() {
	// ---
	// Using new:
	// The built-in new(T) function allocates zeroed memory for type T and returns a pointer to it (*T).
	// It does NOT initialize any internal data structures (e.g., slice backing array, map hash table, channel buffer).
	// Use new when you need a pointer to a value, especially for basic types or structs.
	p := new(int)                        // Allocates memory for an int, returns *int (pointer to zero value)
	fmt.Println("new(int):", *p)         // prints 0 (the zero value for int)
	*p = 42                              // You can assign to the dereferenced pointer
	fmt.Println("after assignment:", *p) // prints 42

	// ---
	// Using make:
	// The built-in make(T, ...) function is used ONLY for slices, maps, and channels.
	// It allocates and initializes the internal data structures needed for these types.
	// For slices: allocates the backing array and returns a slice referencing it.
	// For maps: allocates the hash table and returns a ready-to-use map.
	// For channels: allocates the channel buffer and returns a usable channel.
	s := make([]int, 3)               // Creates a slice of length 3 (all elements zero)
	fmt.Println("make([]int, 3):", s) // prints [0 0 0]

	m := make(map[string]int)               // Creates an empty map ready for use
	m["foo"] = 1                            // Assign a value to a key
	fmt.Println("make(map[string]int):", m) // prints map[foo:1]

	ch := make(chan int, 2)                       // Creates a buffered channel with capacity 2
	ch <- 10                                      // Send value into channel
	ch <- 20                                      // Send another value
	fmt.Println("make(chan int, 2):", <-ch, <-ch) // prints 10 20 (receives values)

	// ---
	// Idiomatic usage and Effective Go guidance:
	// - Use new for pointers to zeroed values (rarely needed for slices, maps, channels)
	// - Use make for slices, maps, channels to initialize their internal data structures
	// - Do NOT use new for slices, maps, or channels (it only gives a pointer to a zero value, which is not usable)
	//   Example: m := new(map[string]int) // m is *map, but *m is nil and cannot be used until assigned with make
	// - Prefer make for slices, maps, channels; use new for other types if you need a pointer

	// ---
	// Summary:
	// new(T) -> *T (pointer to zero value)
	// make(T, ...) -> initialized T (slice, map, or channel, ready to use)
	// See: https://go.dev/doc/effective_go#allocation-new-AND-make
}
