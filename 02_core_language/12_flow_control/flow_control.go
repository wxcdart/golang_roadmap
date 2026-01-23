package main

import "fmt"

// Demonstrates if/else, switch, for (all variants), range, break/continue/return, and labels.

func main() {
	fmt.Println("-- If / Else --")
	x := 7
	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else if x%3 == 0 {
		fmt.Println(x, "is divisible by 3")
	} else {
		fmt.Println(x, "is odd and not divisible by 3")
	}

	fmt.Println("\n-- Switch --")
	switch x {
	case 0:
		fmt.Println("zero")
	case 1, 2, 3:
		fmt.Println("small number")
	default:
		fmt.Println("larger number")
	}

	// switch with initializer and fallthrough example
	switch v := x % 4; v {
	case 0:
		fmt.Println("x % 4 == 0")
	case 3:
		fmt.Println("x % 4 == 3 (fallthrough demonstration)")
		fallthrough
	case 1:
		fmt.Println("handled after fallthrough or case 1")
	}

	fmt.Println("\n-- For loops --")
	// classic for
	for i := 0; i < 3; i++ {
		fmt.Println("i=", i)
	}

	// for as while
	n := 3
	for n > 0 {
		fmt.Println("n=", n)
		n--
	}

	// infinite loop with break
	c := 0
	for {
		c++
		if c > 2 {
			break
		}
		fmt.Println("loop c=", c)
	}

	fmt.Println("\n-- Range over collections --")
	nums := []int{10, 20, 30}
	for idx, val := range nums {
		fmt.Printf("idx=%d val=%d\n", idx, val)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Printf("key=%s value=%d\n", k, v)
	}

	// iterate over string (runes)
	for i, r := range "hi" {
		fmt.Printf("i=%d rune=%c\n", i, r)
	}

	fmt.Println("\n-- break, continue, return --")
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			fmt.Println("continue at", i)
			continue
		}
		if i == 3 {
			fmt.Println("breaking at", i)
			break
		}
		fmt.Println("seen", i)
	}

	fmt.Println("\n-- Labels (use sparingly) --")
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i*j > 1 {
				fmt.Println("breaking out to Outer from", i, j)
				break Outer
			}
			fmt.Println("i,j", i, j)
		}
	}

	fmt.Println("\n-- return from helper to show flow control --")
	helperReturnExample()
}

func helperReturnExample() {
	for i := 0; i < 3; i++ {
		if i == 1 {
			fmt.Println("returning early from helper at", i)
			return
		}
		fmt.Println("helper i", i)
	}
	fmt.Println("this line is unreachable when returned")
}
