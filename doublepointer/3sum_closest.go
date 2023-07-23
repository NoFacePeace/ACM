package doublepointer

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	ans := math.MaxInt
	max := math.MaxInt
	for i := 0; i < n; i++ {
		k := n - 1
		for j := i + 1; j < n; j++ {
			if j == k {
				break
			}
			val := nums[i] + nums[j] + nums[k]
			dst := abs(val, target)
			if max > dst {
				ans = val
				max = dst
			}
			if j < k-1 && val > target {
				k--
				j--
			}
		}
	}
	return ans
}
