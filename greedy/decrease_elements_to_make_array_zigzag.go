package greedy

import "math"

func movesToMakeZigzag(nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	cnt1 := 0
	for i := 1; i < len(nums); i += 2 {
		val := nums[i-1]
		if i+1 < len(nums) {
			val = min(val, nums[i+1])
		}
		if nums[i] < val {
			continue
		}
		cnt1 += nums[i] - val + 1
	}
	cnt2 := 0
	for i := 0; i < len(nums); i += 2 {
		val := math.MaxInt
		if i > 0 {
			val = min(val, nums[i-1])
		}
		if i < len(nums)-1 {
			val = min(val, nums[i+1])
		}
		if nums[i] < val {
			continue
		}
		cnt2 += nums[i] - val + 1
	}
	return min(cnt1, cnt2)
}
