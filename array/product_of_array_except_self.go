package array

func productExceptSelf(nums []int) []int {
	l := len(nums)
	ans := make([]int, l)
	tmp := 1
	for i := 0; i < l; i++ {
		ans[i] = tmp
		tmp *= nums[i]
	}
	tmp = 1
	for i := l - 1; i >= 0; i-- {
		ans[i] *= tmp
		tmp *= nums[i]
	}
	return ans
}
