package str

func evaluate(s string, knowledge [][]string) string {
	m := map[string]string{}
	for _, v := range knowledge {
		m[v[0]] = v[1]
	}
	ans := ""
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			ans += word
			word = ""
			continue
		}
		if s[i] == ')' {
			if _, ok := m[word]; ok {
				ans += m[word]
			} else {
				ans += "?"
			}
			word = ""
			continue
		}
		word += string(s[i])
	}
	ans += word
	return ans
}
