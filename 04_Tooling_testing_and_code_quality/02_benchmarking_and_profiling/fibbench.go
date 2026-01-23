package bench

// ExpensiveFib computes Fibonacci recursively (intentionally slow for demo).
func ExpensiveFib(n int) int {
	if n <= 1 {
		return n
	}
	return ExpensiveFib(n-1) + ExpensiveFib(n-2)
}
