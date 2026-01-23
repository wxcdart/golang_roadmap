// arrays_slices_maps.go
// Demonstrates arrays, slices, 2D slices, and maps in Go with idiomatic usage and comments.
// Includes examples and explanations per Effective Go.

package main

import (
	"fmt"
)

func main() {
	// --- Arrays ---
	// Fixed-size, value type, rarely used directly except for low-level code
	// Go arrays are similar to Java arrays (fixed size, type-safe), but less flexible than Python lists or NumPy arrays.
	// In Python: arr = [1, 2, 3] (dynamic list), or arr = np.array([1,2,3]) for fixed-size NumPy array
	// In Java: int[] arr = {1, 2, 3};
	var arr [3]int // array of 3 ints, zero-initialized
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("Array:", arr)

	// Arrays are value types: assignment copies the entire array.
	arrCopy := arr
	arrCopy[0] = 99
	fmt.Println("After copying and modifying arrCopy:", arrCopy)
	fmt.Println("Original arr remains unchanged (value type):", arr)

	// --- Slices ---
	// Dynamic, flexible view over an array; idiomatic for most use cases
	// Go slices are similar to Python lists (dynamic, resizable) and Java's ArrayList (dynamic, type-safe).
	// In Python: s = [10, 20, 30]
	// In Java: ArrayList<Integer> s = new ArrayList<>(Arrays.asList(10, 20, 30));
	s := []int{10, 20, 30}
	fmt.Println("Slice:", s)

	// Slices are reference types: they describe a segment of an underlying array.
	// Assigning a slice value copies the slice header (pointer, len, cap), not the underlying array.
	sRef := s
	sRef[0] = 111
	fmt.Println("After modifying sRef[0]:", sRef)
	fmt.Println("Original slice s reflects change (shared underlying array):", s)

	// Slicing an array
	// Go slices support slicing syntax like Python (but not as flexible as NumPy)
	// In Python: s2 = arr[1:]
	// In Java: Arrays.copyOfRange(arr, 1, arr.length)
	s2 := arr[1:]
	fmt.Println("Sliced array:", s2)

	// --- 2D Slices (slice of slices) ---
	// Used for matrices, grids, etc.
	// Go 2D slices are similar to Python lists of lists, or NumPy arrays for matrix math.
	// In Python: matrix = [[0,0,0],[0,0,0]]; matrix[0][1] = 7
	// In Java: int[][] matrix = new int[2][3]; matrix[0][1] = 7;
	// For advanced math, Python's NumPy is more powerful; Go requires manual iteration or external packages.
	matrix := make([][]int, 2) // 2 rows
	for i := range matrix {
		matrix[i] = make([]int, 3) // 3 columns per row
	}
	matrix[0][1] = 7
	matrix[1][2] = 9
	fmt.Println("2D Slice (matrix):", matrix)

	// --- Math with 2D Slices (matrix) ---
	// Sum all elements
	total := 0
	for i := range matrix {
		for j := range matrix[i] {
			total += matrix[i][j]
		}
	}
	fmt.Println("Sum of all elements in matrix:", total)

	// Row sums
	for i := range matrix {
		rowSum := 0
		for j := range matrix[i] {
			rowSum += matrix[i][j]
		}
		fmt.Printf("Sum of row %d: %d\n", i, rowSum)
	}

	// Column sums
	if len(matrix) > 0 {
		for col := 0; col < len(matrix[0]); col++ {
			colSum := 0
			for row := 0; row < len(matrix); row++ {
				colSum += matrix[row][col]
			}
			fmt.Printf("Sum of column %d: %d\n", col, colSum)
		}
	}

	// --- Maps ---
	// Key-value store, unordered, dynamic
	// Go maps are similar to Python dicts and Java HashMap.
	// In Python: m = {'apple': 5, 'banana': 3}
	// In Java: Map<String, Integer> m = new HashMap<>(); m.put("apple", 5);
	m := make(map[string]int)
	m["apple"] = 5
	m["banana"] = 3
	fmt.Println("Map:", m)

	// Map lookup and existence check
	v, ok := m["apple"]
	fmt.Printf("m[\"apple\"] = %d, exists? %v\n", v, ok)

	// Iterating over a map (order is not guaranteed)
	for k, v := range m {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	// Map operations: delete, zero-value lookup
	delete(m, "banana")
	v2, ok2 := m["banana"]
	fmt.Printf("After delete: banana exists? %v, value (zero if not): %d\n", ok2, v2)

	// --- append() and capacity behavior ---
	fmt.Println("\n-- append and capacity behavior --")
	small := make([]int, 0, 2) // length 0, capacity 2
	fmt.Printf("before: len=%d cap=%d\n", len(small), cap(small))
	small = append(small, 1)
	// take address of first element to observe possible reallocation later
	p0 := &small[0]
	fmt.Printf("after 1 append: len=%d cap=%d ptr=%p\n", len(small), cap(small), p0)
	small = append(small, 2)
	fmt.Printf("after 2 append: len=%d cap=%d ptr=%p\n", len(small), cap(small), &small[0])
	// appending beyond capacity will cause Go to allocate a new underlying array
	small = append(small, 3)
	fmt.Printf("after 3rd append (beyond cap): len=%d cap=%d ptr=%p\n", len(small), cap(small), &small[0])

	fmt.Println("\n-- Summary: arrays vs slices vs maps --")
	fmt.Println("- Arrays: fixed-length, value types (assign copies the array).")
	fmt.Println("- Slices: dynamic, reference-like headers pointing to an underlying array; use append to grow; capacity determines when Go reallocates.")
	fmt.Println("- Maps: dynamic key-value stores; use delete and check existence with 'v, ok := m[key]'.")

	// --- Effective Go tips ---
	// - Prefer slices over arrays for most cases
	// - Use make to create slices, 2D slices, and maps
	// - Check existence with 'v, ok := m[key]'
	// - Use range for iteration

	// --- Cross-language summary ---
	// Arrays: Go (fixed size, value type), Python (list, dynamic), Java (fixed size)
	// Slices: Go (dynamic, idiomatic), Python (list), Java (ArrayList)
	// 2D Slices: Go (slice of slices), Python (list of lists, NumPy array), Java (2D array)
	// Maps: Go (map), Python (dict), Java (HashMap)
}
