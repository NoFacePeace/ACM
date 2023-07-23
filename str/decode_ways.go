package str

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, l+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= l; i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		if s[i-2] == '2' && s[i-1] <= '6' && s[i-1] >= '0' {
			dp[i] += dp[i-2]
		}
		if s[i-2] == '1' {
			dp[i] += dp[i-2]
		}
		if dp[i] == 0 {
			return 0
		}
	}
	return dp[l]
}
