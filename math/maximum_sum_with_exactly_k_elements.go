package math

func maximizeSum(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	max *= k
	max += k * (k - 1) / 2
	return max
}
