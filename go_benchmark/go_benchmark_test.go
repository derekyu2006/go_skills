package go_benchmark

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sum(1, 2)
	}
}
