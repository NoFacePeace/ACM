package str

import "strings"

func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	m := map[byte]int{}
	for k := range s {
		m[s[k]]++
	}
	var split byte
	for c, v := range m {
		if v < k {
			split = c
			break
		}
	}
	if split == 0 {
		return len(s)
	}
	ans := 0
	for _, v := range strings.Split(s, string(split)) {
		ans = max(ans, longestSubstring(v, k))
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
