package str

func camelMatch(queries []string, pattern string) []bool {
	ans := []bool{}
	for _, v := range queries {
		i := 0
		j := 0
		match := true
		for i < len(v) && j < len(pattern) {
			if v[i] == pattern[j] {
				i++
				j++
				continue
			}
			if v[i] >= 'A' && v[i] <= 'Z' {
				match = false
				break
			}
			i++
		}
		for i < len(v) {
			if v[i] >= 'A' && v[i] <= 'Z' {
				match = false
				break
			}
			i++
		}
		if j < len(pattern) {
			match = false
		}
		ans = append(ans, match)
	}
	return ans
}
