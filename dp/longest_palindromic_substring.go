package dp

func longestPalindrome1(s string) string {
	l := len(s)
	if l == 0 {
		return s
	}
	ans := s[0:1]
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if j+i >= l {
				break
			}
			if j+i == j {
				dp[j][j+i] = true
				continue
			}
			if i == 1 {
				if s[j] == s[j+i] {
					dp[j][j+i] = true
					if len(ans) < i+1 {
						ans = s[j : j+i+1]
					}
				} else {
					dp[j][j+i] = false
				}
				continue
			}
			if dp[j+1][j+i-1] && s[j] == s[j+i] {
				dp[j][j+i] = true
				if len(ans) < i+1 {
					ans = s[j : j+i+1]
				}
			} else {
				dp[j][j+i] = false
			}
		}
	}
	return ans
}
