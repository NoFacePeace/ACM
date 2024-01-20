package stack

func sumSubarrayMins(arr []int) int {
	mod := int(1e9) + 7
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	stack := []int{}
	for k, v := range arr {
		l := len(stack)
		for l > 0 && v <= arr[stack[l-1]] {
			stack = stack[:l-1]
			l--
		}
		l = len(stack)
		if l == 0 {
			left[k] = k + 1
		} else {
			left[k] = k - stack[l-1]
		}
		stack = append(stack, k)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		l := len(stack)
		for l > 0 && arr[i] < arr[stack[l-1]] {
			stack = stack[:l-1]
			l--
		}
		l = len(stack)
		if l == 0 {
			right[i] = n - i
		} else {
			right[i] = stack[l-1] - i
		}
		stack = append(stack, i)
	}
	ans := 0
	for k, x := range arr {
		ans = (ans + left[k]*right[k]*x) % mod
	}
	return ans
}
