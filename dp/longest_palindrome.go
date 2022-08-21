package dp

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return s
	}
	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		}
		return s[0:1]
	}
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	max := 0
	ans := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			k := j + i
			if k >= len(s) {
				continue
			}
			if s[k] != s[j] {
				continue
			}
			if k-j < 3 {
				dp[j][k] = 1
			} else {
				dp[j][k] = dp[j+1][k-1]
			}
			if dp[j][k] == 1 && k-j >= max {
				max = k - j
				ans = s[j : k+1]
			}
		}
	}
	return ans
}
