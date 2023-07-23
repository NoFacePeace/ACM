package dp

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	arr := make([][]byte, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				arr[i][j] = text1[i]
				dp[i+1][j+1] = dp[i][j] + 1
				continue
			}
			dp[i+1][j+1] = dp[i][j+1]
			if dp[i+1][j+1] < dp[i+1][j] {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}
	fmt.Println(arr)
	return dp[m][n]
}
