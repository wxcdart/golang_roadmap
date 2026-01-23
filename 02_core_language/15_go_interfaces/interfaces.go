package main

import "fmt"

// Interfaces in Go are about behavior, not inheritance.
// Implementation is implicit — a type implements an interface by providing its methods.

type Speaker interface {
	Speak() string
}

// Value receiver: both T and *T implement Speaker.
type Dog struct{ Name string }

func (d Dog) Speak() string { return d.Name + " says woof" }

// Pointer receiver: only *T implements Speaker, not T.
type Cat struct{ Name string }

func (c *Cat) Speak() string { return c.Name + " says meow" }

func describe(s Speaker) {
	fmt.Println("describe:", s.Speak())
}

func main() {
	// Implicit implementation
	d := Dog{"Fido"}
	describe(d) // Dog has value receiver; value implements Speaker

	c := Cat{"Whiskers"}
	// describe(c) // compile error: Cat (non-pointer) does NOT implement Speaker
	describe(&c) // pass pointer

	// Slice of interface values (polymorphism)
	animals := []Speaker{d, &c}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}

	// The empty interface (interface{}) can hold any value
	var any interface{}
	any = 42
	fmt.Printf("any=%v (type %T)\n", any, any)
	any = "hello"
	fmt.Printf("any=%v (type %T)\n", any, any)

	// Type assertion: retrieve concrete type from interface value
	if s, ok := any.(string); ok {
		fmt.Println("asserted string value:", s)
	}

	// Type switch: branch on underlying type
	switch v := any.(type) {
	case int:
		fmt.Println("int in type switch:", v)
	case string:
		fmt.Println("string in type switch:", v)
	default:
		fmt.Printf("other type: %T\n", v)
	}

	// Interfaces encourage decoupling: code depends on behavior (interfaces), not concrete types.
	fmt.Println("Go interfaces encourage clean, decoupled, composable code.")
}
