package pointer

func nthUglyNumber(n int) int {
	dp := []int{1}
	n2, n3, n5 := 0, 0, 0
	i := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i < n {
		mn := min(dp[n2]*2, min(dp[n3]*3, dp[n5]*5))
		dp = append(dp, mn)
		if dp[n2]*2 == mn {
			n2++
		}
		if dp[n3]*3 == mn {
			n3++
		}
		if dp[n5]*5 == mn {
			n5++
		}
		i++
	}
	return dp[n-1]
}
