package stack

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	stack := make([]byte, len(s))
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	index := -1
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			index++
			stack[index] = s[i]
			continue
		}
		if index == -1 {
			return false
		}
		if m[stack[index]] == s[i] {
			index--
			continue
		}
		return false
	}
	if index == -1 {
		return true
	}
	return false
}
