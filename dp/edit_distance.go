package dp

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = dp[i][j] + 1
			}
			if dp[i+1][j+1] > dp[i+1][j]+1 {
				dp[i+1][j+1] = dp[i+1][j] + 1
			}
			if dp[i+1][j+1] > dp[i][j+1]+1 {
				dp[i+1][j+1] = dp[i][j+1] + 1
			}
		}
	}
	return dp[m][n]
}
