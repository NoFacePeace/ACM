package str

func calculate(s string) int {
	operator := '+'
	num := 0
	stack := []int{}
	for k, v := range s {
		isDigit := v >= '0' && v <= '9'
		if isDigit {
			num = num*10 + int(v-'0')
		}
		if !isDigit && v != ' ' || k == len(s)-1 {
			switch operator {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			num = 0
			operator = v
		}
	}
	ans := 0
	for _, v := range stack {
		ans += v
	}
	return ans
}
