package array

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Slice(coins, func(a, b int) bool {
		return coins[a] < coins[b]
	})
	target := 0
	for _, v := range coins {
		if v <= target+1 {
			target += v
		}
	}
	return target + 1
}
