package stack

func backspaceCompare(s string, t string) bool {
	stack1 := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			stack1 = append(stack1, s[i])
			continue
		}
		if len(stack1) == 0 {
			continue
		}
		stack1 = stack1[:len(stack1)-1]
	}
	stack2 := []byte{}
	for i := 0; i < len(t); i++ {
		if t[i] != '#' {
			stack2 = append(stack2, t[i])
			continue
		}
		if len(stack2) == 0 {
			continue
		}
		stack2 = stack2[:len(stack2)-1]
	}
	if len(stack1) != len(stack2) {
		return false
	}
	for i := 0; i < len(stack1); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}
	return true
}
