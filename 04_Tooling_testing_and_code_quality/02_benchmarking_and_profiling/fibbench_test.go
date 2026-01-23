package bench

import (
	"testing"
)

func BenchmarkExpensiveFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ExpensiveFib(20) // moderate work
	}
}
