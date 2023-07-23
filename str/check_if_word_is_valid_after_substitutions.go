package str

func isValid(s string) bool {
	i := 0
	for i < len(s)-2 && len(s) != 0 {
		if s[i:i+3] == "abc" {
			s = s[:i] + s[i+3:]
			i -= 2
			if i < 0 {
				i = 0
			}
			continue
		}
		i++
	}
	if s == "" {
		return true
	}
	return false
}
