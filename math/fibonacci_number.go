package math

func fib(n int) int {
	var f func(n int) int
	f = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		return f(n-1) + f(n-2)
	}
	return f(n)
}
