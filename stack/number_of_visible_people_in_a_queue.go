package stack

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	ans := make([]int, n)
	if n == 0 {
		return ans
	}
	stack := []int{}
	stack = append(stack, heights[n-1])
	for i := n - 2; i >= 0; i-- {
		cnt := 0
		l := len(stack)
		for j := l - 1; j >= 0; j-- {
			if heights[i] <= stack[j] {
				cnt++
				break
			}
			cnt++
		}
		ans[i] = cnt
		for len(stack) > 0 {
			l := len(stack)
			if stack[l-1] > heights[i] {
				break
			}
			stack = stack[:l-1]
		}
		stack = append(stack, heights[i])
	}
	return ans
}
