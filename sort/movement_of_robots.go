package sort

import "sort"

func sumDistance(nums []int, s string, d int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	for i := 0; i < n; i++ {
		if s[i] == 'L' {
			nums[i] -= d
		} else {
			nums[i] += d
		}
	}
	mod := int(1e9) + 7
	sort.Ints(nums)
	ans := 0
	prev := 0
	for i := 1; i < n; i++ {
		prev += (nums[i] - nums[i-1]) * i
		ans += prev
		ans %= mod
	}
	return ans
}
