package str

func isAnagram(s string, t string) bool {
	m1 := map[byte]int{}
	for k := range s {
		m1[s[k]] += 1
	}
	m2 := map[byte]int{}
	for k := range t {
		m2[t[k]] += 1
	}
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	for k := range m2 {
		if m2[k] != m1[k] {
			return false
		}
	}
	return true
}
