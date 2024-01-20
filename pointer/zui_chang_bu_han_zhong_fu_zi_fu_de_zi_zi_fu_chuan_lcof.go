package pointer

func lengthOfLongestSubstring(s string) int {
	m := map[byte]bool{}
	n := len(s)
	i := 0
	max := 0
	for j := 0; j < n; j++ {
		if !m[s[j]] {
			m[s[j]] = true
			if j-i+1 > max {
				max = j - i + 1
			}
			continue
		}
		for s[i] != s[j] {
			m[s[i]] = false
			i++
		}
		i++
	}
	if n-i > max {
		max = n - i
	}
	return max
}
