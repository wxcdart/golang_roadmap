package main

import "fmt"

// Demonstrates primitive types, type inference, zero values, constants/iota,
// explicit conversions, and formatted printing with fmt.Printf.

const (
	// iota generates successive untyped integer constants starting at 0.
	First = iota
	Second
	Third
)

func main() {
	// Primitive types
	var s string = "hello"
	var i int = 10
	var f float64 = 3.1415
	var b bool = true

	// Type inference using :=
	inferredStr := "inferred string"
	inferredNum := 42

	// Zero values
	var zeroInt int
	var zeroString string
	var zeroBool bool
	var zeroFloat float64

	fmt.Println("-- Basic types and inference --")
	fmt.Printf("s (%T) = %q\n", s, s)
	fmt.Printf("i (%T) = %d\n", i, i)
	fmt.Printf("f (%T) = %f\n", f, f)
	fmt.Printf("b (%T) = %t\n", b, b)
	fmt.Printf("inferredStr (%T) = %q\n", inferredStr, inferredStr)
	fmt.Printf("inferredNum (%T) = %d\n", inferredNum, inferredNum)

	fmt.Println("\n-- Zero values --")
	fmt.Printf("zeroInt (%T) = %d\n", zeroInt, zeroInt)
	fmt.Printf("zeroString (%T) = %q\n", zeroString, zeroString)
	fmt.Printf("zeroBool (%T) = %t\n", zeroBool, zeroBool)
	fmt.Printf("zeroFloat (%T) = %f\n", zeroFloat, zeroFloat)

	fmt.Println("\n-- Constants and iota --")
	fmt.Printf("First=%d, Second=%d, Third=%d\n", First, Second, Third)

	// Explicit conversion
	var myInt int = 7
	var myFloat float64 = 2.5
	sum := float64(myInt) + myFloat
	fmt.Printf("Explicit conversion: float64(myInt) + myFloat = %f\n", sum)

	// Formatting examples with verbs
	fmt.Println("\n-- Formatting examples --")
	fmt.Printf("%%v (default) myInt=%v myFloat=%v\n", myInt, myFloat)
	fmt.Printf("%%T (type) myInt=%T myFloat=%T\n", myInt, myFloat)
	fmt.Printf("%%q (quoted) s=%q\n", s)
}
