package dp

func minPathCost(grid [][]int, moveCost [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			for k := 0; k < n; k++ {
				if k == 0 {
					dp[i][j] = dp[i-1][k] + moveCost[grid[i-1][k]][j]
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+moveCost[grid[i-1][k]][j])
			}
			dp[i][j] += grid[i][j]
		}

	}
	mn := dp[m-1][n-1]
	for i := 0; i < n; i++ {
		if dp[m-1][i] < mn {
			mn = dp[m-1][i]
		}
	}
	return mn
}
