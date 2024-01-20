package stack

import "math"

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	for i := 0; i < n/2; i++ {
		postorder[i], postorder[n-1-i] = postorder[n-1-i], postorder[i]
	}
	stack := []int{}
	max := math.MaxInt
	for _, v := range postorder {
		if v >= max {
			return false
		}
		for len(stack) > 0 {
			n := len(stack)
			if stack[n-1] < v {
				break
			}
			max = stack[n-1]
			stack = stack[:n-1]
		}
		stack = append(stack, v)
	}
	return true
}
