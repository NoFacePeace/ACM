package array

func longestWPI(hours []int) int {
	l := len(hours)
	if l <= 0 {
		return 0
	}
	arr := make([]int, l+1)
	stack := []int{0}
	for i := 1; i <= l; i++ {
		if hours[i-1] > 8 {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = arr[i-1] - 1
		}
		top := stack[len(stack)-1]
		if arr[i] < arr[top] {
			stack = append(stack, i)
		}
	}
	ans := 0
	for i := l; i >= 0; i-- {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if arr[i] <= arr[top] {
				break
			}
			if i-top > ans {
				ans = i - top
			}
			stack = stack[:len(stack)-1]
		}

	}
	return ans
}
