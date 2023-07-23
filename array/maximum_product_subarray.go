package array

func maxProduct(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	ans := nums[0]
	min := nums[0]
	max := nums[0]
	smallest := func(a, b, c int) int {
		smaller := func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}
		return smaller(smaller(a, b), c)
	}
	biggest := func(a, b, c int) int {
		bigger := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}
		return bigger(bigger(a, b), c)
	}
	for i := 1; i < l; i++ {
		min, max = smallest(min*nums[i], nums[i], max*nums[i]), biggest(min*nums[i], nums[i], max*nums[i])
		if max > ans {
			ans = max
		}
	}
	return ans
}
