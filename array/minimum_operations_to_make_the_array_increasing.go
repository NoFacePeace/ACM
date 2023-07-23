package array

func minOperations2(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if nums[i] > nums[i-1] {
			continue
		}
		ans += nums[i-1] - nums[i] + 1
		nums[i] = nums[i-1] + 1
	}
	return ans
}
