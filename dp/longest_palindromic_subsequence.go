package dp

func longestPalindromeSubseq(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	max := 1
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if j+i >= l {
				break
			}
			if j+i == j {
				dp[j][j+i] = 1
				continue
			}
			if i == 1 {
				if s[j] == s[j+i] {
					dp[j][j+i] = 2
					if max < 2 {
						max = 2
					}
				} else {
					dp[j][j+i] = 1
				}
				continue
			}
			if s[j] != s[j+i] {
				dp[j][j+i] = dp[j][j+i-1]
				if dp[j][j+i] < dp[j+1][j+i] {
					dp[j][j+i] = dp[j+1][j+i]
				}
				continue
			}
			dp[j][j+i] = dp[j+1][j+i-1] + 2
			if dp[j][j+i] > max {
				max = dp[j][j+i]
			}
		}
	}
	return max
}
