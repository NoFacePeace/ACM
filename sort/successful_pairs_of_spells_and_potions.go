package sort

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := []int{}
	n := len(potions)
	binarySearch := func(num int) int {
		l := 0
		r := n
		for l < r {
			m := (l + r) / 2
			if potions[m]*num < int(success) {
				l = m + 1
			} else {
				r = m
			}
		}
		return r
	}
	for _, v := range spells {
		idx := binarySearch(v)
		ans = append(ans, n-idx)
	}
	return ans
}
