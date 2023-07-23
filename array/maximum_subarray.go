package array

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
