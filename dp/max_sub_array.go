package dp

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if sum < nums[i] {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
