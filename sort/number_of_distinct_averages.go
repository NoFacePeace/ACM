package sort

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	m := map[float64]int{}
	for i := 0; i < n/2; i++ {
		f := float64(nums[i]) + float64(nums[n-i-1])
		f /= 2
		m[f]++
	}
	if n%2 != 0 {
		f := float64(nums[n/2])
		m[f]++
	}
	return len(m)
}
