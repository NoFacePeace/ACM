package str

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return -1
	}
	if len(haystack) < len(needle) {
		return -1
	}
	arr := kmp(needle)
	i := 0
	j := 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
			continue
		}
		if j == 0 {
			i++
			continue
		}
		j = arr[j-1]
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func kmp(s string) []int {
	ans := make([]int, len(s))
	if len(s) == 0 {
		return ans
	}
	if len(s) == 1 {
		return ans
	}
	for i := 1; i < len(s); i++ {
		j := ans[i-1]
		for j > 0 && s[j] != s[i] {
			j = ans[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		ans[i] = j
	}
	return ans
}
