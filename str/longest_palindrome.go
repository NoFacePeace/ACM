package str

func longestPalindrome2(s string) int {
	m := map[byte]int{}
	isOdd := false
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	ans := 0
	for _, v := range m {
		if v%2 == 1 {
			isOdd = true
			ans += v - 1
		} else {
			ans += v
		}
	}
	if isOdd {
		return ans + 1
	}
	return ans
}
