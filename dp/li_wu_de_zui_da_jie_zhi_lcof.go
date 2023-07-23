package dp

func maxValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
				continue
			}
			if i == 0 {
				dp[j] = dp[j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				dp[j] += grid[i][j]
				continue
			}
			dp[j] = max(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}
