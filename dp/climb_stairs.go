package dp

func climbStairs(n int) int {
	first := 0
	second := 1
	for i := 0; i < n; i++ {
		tmp := first
		first = second
		second = tmp + second
	}
	return second
}
