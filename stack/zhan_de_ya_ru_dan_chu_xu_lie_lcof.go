package stack

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	n := len(pushed)
	i := 0
	j := 0
	for i < n && j < n {
		if pushed[i] == popped[j] {
			i++
			j++
			continue
		}
		if len(stack) == 0 {
			stack = append(stack, pushed[i])
			i++
			continue
		}
		if stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
			continue
		}
		stack = append(stack, pushed[i])
		i++
	}
	for j < n {
		n := len(stack)
		if stack[n-1] != popped[j] {
			return false
		}
		j++
		stack = stack[:n-1]
	}
	return true
}
