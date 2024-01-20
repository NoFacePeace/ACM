package greedy

func longestValidParentheses(s string) int {
	left := 0
	right := 0
	max := 0
	bytes := []byte(s)
	n := len(bytes)
	for i := 0; i < n; i++ {
		if bytes[i] == '(' {
			left++
		} else {
			right++
		}
		if left < right {
			left = 0
			right = 0
			continue
		}
		if left == right && left*2 > max {
			max = left * 2
		}
	}
	left = 0
	right = 0
	for i := n - 1; i >= 0; i-- {
		if bytes[i] == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			left = 0
			right = 0
			continue
		}
		if left == right && left*2 > max {
			max = left * 2
		}
	}
	return max
}
