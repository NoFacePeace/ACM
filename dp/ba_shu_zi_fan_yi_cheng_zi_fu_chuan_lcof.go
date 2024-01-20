package dp

import "strconv"

func translateNum(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i+1] = 1
			continue
		}
		if str[i-1] == '1' {
			dp[i+1] += dp[i] + dp[i-1]
			continue
		}
		if str[i-1] == '2' && str[i] <= '5' {
			dp[i+1] += dp[i] + dp[i-1]
			continue
		}
		dp[i+1] = dp[i]
	}
	return dp[n]
}
