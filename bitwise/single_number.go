package bitwise

func singleNumber(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans
}
