package str

func isIsomorphic(s string, t string) bool {
	m := map[byte]byte{}
	l := len(s)
	if l != len(t) {
		return false
	}
	used := map[byte]bool{}
	for i := 0; i < l; i++ {
		if _, ok := m[s[i]]; ok {
			if m[s[i]] != t[i] {
				return false
			}
			continue
		}
		if used[t[i]] {
			return false
		}
		m[s[i]] = t[i]
		used[t[i]] = true
	}
	return true
}
