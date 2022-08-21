package slidwindow

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	m := map[byte]bool{}
	m[s[0]] = true
	max := 0
	left := 0
	right := 1
	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			delete(m, s[left])
			left++
		} else {
			m[s[right]] = true
			right++
		}
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}
