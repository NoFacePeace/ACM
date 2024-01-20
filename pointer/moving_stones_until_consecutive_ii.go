package pointer

import "sort"

func numMovesStonesII(stones []int) []int {
	sort.Ints(stones)
	l := len(stones)
	if stones[l-1]-stones[0]+1 == l {
		return []int{0, 0}
	}
	right := stones[l-2] - stones[0] + 1 - (l - 1)
	left := stones[l-1] - stones[1] + 1 - (l - 1)
	var max int
	if right > left {
		max = right
	} else {
		max = left
	}
	min := l
	j := 0
	for i := 0; i < l; i++ {
		for j+1 < l && stones[j+1]-stones[i]+1 <= l {
			j++
		}
		if j-i+1 == l-1 && stones[j]-stones[i]+1 == l-1 {
			if min > 2 {
				min = 2
			}
		} else {
			if min > l-(j-i+1) {
				min = l - (j - i + 1)
			}
		}
	}
	return []int{min, max}
}
