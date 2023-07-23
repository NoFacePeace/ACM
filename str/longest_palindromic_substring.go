package str

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	max := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i+j >= len(s) {
				break
			}
			if i == 0 {
				dp[j][j] = true
				if i >= len(max) {
					max = s[j : j+i+1]
				}
				continue
			}
			if i == 1 {
				if s[j] == s[i+j] {
					dp[j][j+i] = true
					if i >= len(max) {
						max = s[j : j+i+1]
					}
					continue
				}
				dp[j][j+i] = false
				continue
			}
			if s[j] == s[j+i] && dp[j+1][j+i-1] {
				dp[j][j+i] = true
				if i >= len(max) {
					max = s[j : j+i+1]
				}
				continue
			}
			dp[j][j+i] = false
		}
	}
	return max
}

func longestPalindrome1(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	s = t
	arm_len := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var cur_arm_len int
		if right >= i {
			i_sym := j*2 - i
			min_arm_len := min(arm_len[i_sym], right-i)
			cur_arm_len = expand(s, i-min_arm_len, i+min_arm_len)
		} else {
			cur_arm_len = expand(s, i, i)
		}
		arm_len = append(arm_len, cur_arm_len)
		if i+cur_arm_len > right {
			j = i
			right = i + cur_arm_len
		}
		if cur_arm_len*2+1 > end-start {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
