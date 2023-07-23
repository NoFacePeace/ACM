package array

func canJump(nums []int) bool {
	l := len(nums)
	if l == 0 {
		return true
	}
	max := nums[0]
	for i := 1; i < l; i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
}
