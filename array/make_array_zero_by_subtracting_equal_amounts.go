package array

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	reduce := 0
	cnt := 0
	for i := 0; i < len(nums); i++ {
		dist := nums[i] - reduce
		if dist != 0 {
			cnt++
		}
		reduce += dist
	}
	return cnt
}
