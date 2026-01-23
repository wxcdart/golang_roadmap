package mathutil

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Mul returns the product of a and b.
func Mul(a, b int) int { return a * b }

// Fib computes the n-th Fibonacci number (inefficient recursive version,
// included here so tests/benchmarks can exercise a non-trivial function).
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
