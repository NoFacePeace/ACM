package stack

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	i := -1
	stack := make([]byte, len(s))
	for k := range s {
		if _, ok := m[s[k]]; ok {
			i += 1
			stack[i] = s[k]
			continue
		}
		if i == -1 {
			return false
		}
		if s[k] != m[stack[i]] {
			return false
		}
		i -= 1
	}
	return i == -1
}
