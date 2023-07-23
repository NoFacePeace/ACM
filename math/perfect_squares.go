package math

func numSquares(n int) int {
	dp := map[int]int{}

	for i := 1; i <= n; i++ {
		max := i
		for j := 1; j*j <= i; j++ {
			count := dp[i-j*j]
			if count < max {
				max = count
			}
			dp[i] = 1 + max
		}
	}
	return dp[n]
}
