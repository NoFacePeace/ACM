package str

func firstUniqChar(s string) int {
	m := map[byte]int{}
	for k := range s {
		m[s[k]]++
	}
	for k := range s {
		if m[s[k]] == 1 {
			return k
		}
	}
	return -1
}
