package go_benchmark

// 斐波那契数列
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Sum(a, b int) int {
	return a + b
}
