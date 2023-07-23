package sort

import "sort"

func matrixSum(nums [][]int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	m := len(nums[0])
	if m == 0 {
		return 0
	}
	arr := make([]int, m)
	for i := 0; i < n; i++ {
		sort.Ints(nums[i])
		for j := 0; j < m; j++ {
			if nums[i][j] > arr[j] {
				arr[j] = nums[i][j]
			}
		}
	}
	ans := 0
	for _, v := range arr {
		ans += v
	}
	return ans
}
