package other

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	a := 1
	b := 2
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return a
}
