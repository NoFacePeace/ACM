package str

func isAcronym(words []string, s string) bool {
	m := len(words)
	n := len(s)
	if m != n {
		return false
	}
	for i := 0; i < n; i++ {
		if len(words[i]) == 0 {
			return false
		}
		if words[i][0] != s[i] {
			return false
		}
	}
	return true
}
