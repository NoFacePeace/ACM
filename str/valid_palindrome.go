package str

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	ss := ""
	for k := range s {
		if s[k] >= 'a' && s[k] <= 'z' {
			ss += string(s[k])
		}
		if s[k] >= '0' && s[k] <= '9' {
			ss += string(s[k])
		}
		if s[k] >= 'A' && s[k] <= 'Z' {
			ss += string(toLower(s[k]))
		}
	}
	i, j := 0, len(ss)-1
	for i < j {
		if ss[i] == ss[j] {
			i += 1
			j -= 1
			continue
		}
		return false
	}
	return true
}

func toLower(b byte) byte {
	return b + 'a' - 'A'
}
