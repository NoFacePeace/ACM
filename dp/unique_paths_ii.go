package dp

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 && obstacleGrid[i][j] != 1 {
				dp[i][j] = 1
				continue
			}
			if i == 0 && obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i][j-1]
				continue
			}
			if j == 0 && obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j]
				continue
			}
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
