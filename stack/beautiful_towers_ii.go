package stack

func maximumSumOfHeights(maxHeights []int) int64 {
	n := len(maxHeights)
	left := make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		sum := maxHeights[i]
		for len(stack) > 0 {
			l := len(stack)
			if maxHeights[stack[l-1]] < maxHeights[i] {
				break
			}
			stack = stack[:l-1]
		}
		if len(stack) > 0 {
			l := len(stack)
			sum += left[stack[l-1]]
			sum += (i - stack[l-1] - 1) * maxHeights[i]
		} else {
			sum += i * maxHeights[i]
		}
		left[i] = sum
		stack = append(stack, i)
	}
	stack = []int{}
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		sum := maxHeights[i]
		for len(stack) > 0 {
			l := len(stack)
			if maxHeights[stack[l-1]] < maxHeights[i] {
				break
			}
			stack = stack[:l-1]
		}
		if len(stack) > 0 {
			l := len(stack)
			sum += right[stack[l-1]]
			sum += (stack[l-1] - i - 1) * maxHeights[i]
		} else {
			sum += (n - i - 1) * maxHeights[i]
		}
		right[i] = sum
		stack = append(stack, i)
	}
	mx := 0
	for i := 0; i < n; i++ {
		if left[i]+right[i]-maxHeights[i] > mx {
			mx = left[i] + right[i] - maxHeights[i]
		}
	}
	return int64(mx)
}
