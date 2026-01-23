package main

import "fmt"

// Demonstrates structs, embedding, methods with value and pointer receivers,
// and guidance on choosing structs vs maps.

// Person models an entity with state and behavior.
type Person struct {
	Name string
	Age  int
}

// SayHello is a value-receiver method: works on a copy but doesn't modify original.
func (p Person) SayHello() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d", p.Name, p.Age)
}

// Birthday is a pointer-receiver method: modifies the original struct.
func (p *Person) Birthday() {
	p.Age++
}

// Employee composes Person via embedding; Employee gets Person's methods.
type Employee struct {
	Person
	Position string
}

// Promote modifies the Employee (pointer receiver) to change Position.
func (e *Employee) Promote(newTitle string) {
	e.Position = newTitle
}

func main() {
	// Struct initialization (composite literal)
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.SayHello())

	// Value receiver: calling on value copies the struct
	copyP := p
	copyP.Birthday() // modifies copy
	fmt.Println("after copyP.Birthday(): copy age=", copyP.Age, "original age=", p.Age)

	// Pointer receiver modifies original
	p.Birthday()
	fmt.Println("after p.Birthday(): original age=", p.Age)

	// Embedding example
	e := Employee{Person: Person{Name: "Bob", Age: 40}, Position: "Engineer"}
	fmt.Println(e.SayHello()) // Employee inherits Person.SayHello
	e.Promote("Senior Engineer")
	fmt.Println("Employee position:", e.Position)

	// Zero-value struct
	var zero Person
	fmt.Printf("zero-value Person: %+v\n", zero)

	// When to use structs vs maps:
	// - Use structs when modeling typed data with shape and behavior (methods, receivers).
	// - Use maps for flexible key-value lookup where keys are dynamic or unknown.
	// Example: storing many records for fast lookup
	directory := make(map[string]Person)
	directory["alice"] = p
	directory["bob"] = e.Person
	if v, ok := directory["alice"]; ok {
		fmt.Println("lookup alice:", v.SayHello())
	}

	// Maps vs Struct: maps are great for dynamic collections, but lack methods and compile-time shape.
	fmt.Println("Choose structs for shape+behavior; maps for dynamic key/value storage.")
}
