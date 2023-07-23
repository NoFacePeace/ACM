package str

func prefixCount(words []string, pref string) int {
	ans := 0
	for _, s := range words {
		isMatch := true
		for i := 0; i < len(pref); i++ {
			if i >= len(s) {
				isMatch = false
				break
			}
			if s[i] != pref[i] {
				isMatch = false
			}
		}
		if isMatch {
			ans += 1
		}
	}
	return ans
}
