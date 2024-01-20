package dp

func longestValidParentheses(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	max := 0
	for i := 1; i < n; i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i+1] = dp[i-1] + 2

		} else {
			idx := i - dp[i] - 1
			if idx < 0 {
				continue
			}
			if s[idx] == '(' {
				dp[i+1] = dp[i] + 2 + dp[idx]
			}
		}
		if dp[i+1] > max {
			max = dp[i+1]
		}
	}
	return max
}
