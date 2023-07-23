package str

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]string, len(tokens))
	index := 0
	m := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	for i := 0; i < len(tokens); i++ {
		if _, ok := m[tokens[i]]; !ok {
			stack[index] = tokens[i]
			index++
			continue
		}
		num2, _ := strconv.Atoi(stack[index-1])
		num1, _ := strconv.Atoi(stack[index-2])
		switch tokens[i] {
		case "+":
			num1 = num1 + num2
		case "-":
			num1 = num1 - num2
		case "*":
			num1 = num1 * num2
		case "/":
			num1 = num1 / num2
		}
		index--
		stack[index-1] = strconv.Itoa(num1)
	}
	ans, _ := strconv.Atoi(stack[0])
	return ans
}
