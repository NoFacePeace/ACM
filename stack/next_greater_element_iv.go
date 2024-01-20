package stack

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i := 0
	stack := []int{}
	for i < n {
		ans[i] = -1
		l := len(stack)
		if l == 0 {
			stack = append(stack, i)
			i++
			continue
		}
		if nums[i] <= nums[stack[l-1]] {
			stack = append(stack, i)
			i++
			continue
		}
		j := i + 1
		for l > 0 && nums[stack[l-1]] < nums[i] {
			for j < n {
				if nums[j] > nums[stack[l-1]] {
					break
				}
				j++
			}
			if j == n {
				ans[stack[l-1]] = -1
			} else {
				ans[stack[l-1]] = nums[j]
			}
			stack = stack[:l-1]
			l--
		}
	}
	return ans
}
