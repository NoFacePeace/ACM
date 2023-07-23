package array

func rob(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if i == 1 {
			nums[i] = max(nums[i], nums[i-1])
			continue
		}
		nums[i] = max(nums[i-2]+nums[i], nums[i-1])
	}
	return nums[len(nums)-1]
}
