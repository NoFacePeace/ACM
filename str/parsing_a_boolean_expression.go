package str

func parseBoolExpr(expression string) bool {
	var recur func(s string) bool
	recur = func(s string) bool {
		if s == "t" {
			return true
		}
		if s == "f" {
			return false
		}
		l := len(s)
		if s[0] == '!' {
			return !recur(s[2 : l-1])
		}
		arr := []string{}
		stack := []byte{}
		j := 2
		for i := 2; i < l-1; i++ {
			if s[i] == '(' {
				stack = append(stack, '(')
				continue
			}
			if s[i] == ')' {
				stack = stack[1:]
				continue
			}
			if s[i] == ',' && len(stack) > 0 {
				continue
			}
			if s[i] == ',' {
				arr = append(arr, s[j:i])
				j = i + 1
				continue
			}
		}
		arr = append(arr, s[j:l-1])
		for i := 0; i < len(arr); i++ {
			result := recur(arr[i])
			if s[0] == '|' && result {
				return true
			}
			if s[0] == '&' && !result {
				return false
			}
		}
		return s[0] != '|'
	}
	return recur(expression)
}
