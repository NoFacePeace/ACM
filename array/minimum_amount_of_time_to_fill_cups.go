package array

import "sort"

func fillCups(amount []int) int {
	sort.Slice(amount, func(a, b int) bool {
		return amount[a] > amount[b]
	})
	if amount[0] >= amount[1]+amount[2] {
		return amount[0]
	}
	return amount[0] + (amount[1]+amount[2]-amount[0]+1)/2
}
