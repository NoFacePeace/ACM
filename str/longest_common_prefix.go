package str

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minStr := strs[0]
	ans := ""
	for _, v := range strs {
		if len(v) < len(minStr) {
			minStr = v
		}
	}
	for i := range minStr {
		for j := range strs {
			if strs[j][i] != minStr[i] {
				return ans
			}
		}
		ans += string(minStr[i])
	}
	return ans
}
