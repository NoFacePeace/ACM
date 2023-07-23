package math

func tribonacci(n int) int {
	a := 0
	b := 1
	c := 1
	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	if n == 2 {
		return c
	}
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
