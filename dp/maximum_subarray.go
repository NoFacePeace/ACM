package dp

func maxSubArray(nums []int) int {
	max := 0
	sum := 0
	n := len(nums)
	if n == 0 {
		return 0
	}
	sum = nums[0]
	max = nums[0]
	for i := 1; i < n; i++ {
		if sum+nums[i] > nums[i] {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
