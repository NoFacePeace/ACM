package dp

func findBall(grid [][]int) []int {
	m := len(grid)
	if m == 0 {
		return []int{}
	}
	n := len(grid[0])
	if n == 0 {
		return []int{}
	}
	dp := make([][]int, m)
	for i := m - 1; i >= 0; i-- {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {

			if j == 0 && grid[i][j] == -1 {
				dp[i][j] = -1
				continue
			}
			if j == n-1 && grid[i][j] == 1 {
				dp[i][j] = -1
				continue
			}
			if i == m-1 {
				if grid[i][j] == 1 && grid[i][j+1] == 1 {
					dp[i][j] = j + 1
					continue
				}
				if grid[i][j] == -1 && grid[i][j-1] == -1 {
					dp[i][j] = j - 1
					continue
				}
				dp[i][j] = -1
				continue
			}
			if grid[i][j] == 1 && grid[i][j+1] == 1 {
				dp[i][j] = dp[i+1][j+1]
				continue
			}
			if grid[i][j] == -1 && grid[i][j-1] == -1 {
				dp[i][j] = dp[i+1][j-1]
				continue
			}
			dp[i][j] = -1
		}
	}
	return dp[0]
}
