package array

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	count := func(n int) int {
		ret := 0
		cur := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= n {
				cur++
			} else {
				cur = 0
			}
			ret += cur
		}
		return ret
	}
	return count(right) - count(left-1)
}
