package dp

func maxAbsoluteSum(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := 0
	ans := max
	for i := 0; i < n; i++ {
		max = MaxInt(max+nums[i], nums[i])
		ans = MaxInt(max, ans)
		nums[i] = -nums[i]
	}
	max = 0
	for i := 0; i < n; i++ {
		max = MaxInt(max+nums[i], nums[i])
		ans = MaxInt(max, ans)
		nums[i] = -nums[i]
	}
	return ans
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
