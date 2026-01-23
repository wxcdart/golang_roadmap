package main

import (
	"errors"
	"fmt"
)

// Demonstrates: function definitions, multiple returns, variadic params,
// anonymous functions, and passing functions as arguments.

// add returns a sum and nil error.
func add(a, b int) (int, error) {
	return a + b, nil
}

// div returns quotient and error when dividing by zero.
func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// variadicSum accepts any count of ints.
func variadicSum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// applyFunc demonstrates passing a function as an argument.
func applyFunc(a int, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func main() {
	// Basic function call and multiple return values
	s, err := add(2, 3)
	if err != nil {
		fmt.Println("add error:", err)
	} else {
		fmt.Println("add(2,3) =", s)
	}

	// Division with error handling
	if q, err := div(10, 2); err == nil {
		fmt.Println("div(10,2) =", q)
	}
	if _, err := div(1, 0); err != nil {
		fmt.Println("div error (expected):", err)
	}

	// Variadic parameters and type inference
	fmt.Println("variadicSum(1,2,3,4) =", variadicSum(1, 2, 3, 4))
	nums := []int{5, 6, 7}
	fmt.Println("variadicSum(nums...) =", variadicSum(nums...))

	// Anonymous function assigned to a variable
	multiply := func(a, b int) int { return a * b }
	fmt.Println("multiply(3,4) =", multiply(3, 4))

	// Passing function as argument (higher-order)
	result := applyFunc(4, 5, func(x, y int) int { return x - y })
	fmt.Println("applyFunc with anonymous subtract =", result)

	// Using a named function passed as argument
	sumByApply := applyFunc(7, 8, func(x, y int) int { return x + y })
	fmt.Println("applyFunc with add-like =", sumByApply)

	// Returning multiple values and using blank identifier to ignore one
	if v, _ := add(10, 20); true {
		fmt.Println("add(10,20) (ignored err) =", v)
	}
}
