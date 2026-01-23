package main

import "fmt"

// Demonstrates pointers, dereferencing, nil checks, value vs reference semantics,
// and a short note about Go's garbage collector.

type Node struct {
	Val  int
	Next *Node
}

func incVal(v int)  { v++ }
func incPtr(v *int) { (*v)++ }

func main() {
	// Basic pointer usage
	a := 5
	p := &a // p holds the address of a
	fmt.Printf("a=%d p=%p *p=%d\n", a, p, *p)

	// Modify through pointer (dereference)
	*p = 10
	fmt.Println("after *p=10 a=", a)

	// Nil pointer check
	var q *int
	if q == nil {
		fmt.Println("q is nil; don't dereference")
	}

	// new returns a pointer to a zeroed value
	pn := new(int)
	fmt.Printf("pn=%p *pn=%d\n", pn, *pn)
	*pn = 7
	fmt.Println("after *pn=7 *pn=", *pn)

	// Value vs reference semantics
	x := 1
	incVal(x)
	fmt.Println("after incVal(x):", x)
	incPtr(&x)
	fmt.Println("after incPtr(&x):", x)

	// Pointers to structs and linked-like structure
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n1.Next = n2
	fmt.Printf("n1=%p n2=%p n1.Next=%p n1.Next.Val=%d\n", n1, n2, n1.Next, n1.Next.Val)

	// Arrays are value types (copy), slices are reference-like (share underlying array)
	arr := [2]int{1, 2}
	arr2 := arr
	arr2[0] = 9
	fmt.Println("arr (value):", arr, "arr2 (copy):", arr2)

	sl := []int{1, 2}
	sl2 := sl
	sl2[0] = 9
	fmt.Println("sl (shared underlying array):", sl, "sl2:", sl2)

	// Addresses of elements
	fmt.Printf("&arr[0]=%p &sl[0]=%p\n", &arr[0], &sl[0])

	// Avoid dereferencing nil pointers; GC note
	fmt.Println("Go's garbage collector frees memory for unreachable objects automatically.")
}
