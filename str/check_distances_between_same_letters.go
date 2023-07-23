package str

func checkDistances(s string, distance []int) bool {
	m := map[byte]bool{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		m[s[i]] = true
		dist := distance[int(s[i]-'a')] + 1
		if i+dist >= len(s) {
			return false
		}
		if s[i+dist] != s[i] {
			return false
		}
	}
	return true
}
