package dp

func cuttingRope(n int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, n)
	var f func(n int) int
	f = func(n int) int {
		if n == 1 {
			return 1
		}
		if dp[n] != 0 {
			return dp[n]
		}
		dp[n] = n
		for i := 1; i < n; i++ {
			dp[n] = max(dp[n], f(i)*f(n-i))
		}
		return dp[n]
	}
	mx := 0
	for i := 1; i < n; i++ {
		mx = max(mx, f(i)*f(n-i))
	}
	return mx
}
