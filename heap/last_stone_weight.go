package heap

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	i := len(stones) - 1
	for i >= 1 {
		if stones[i] == stones[i-1] {
			stones[i] = 0
			stones[i-1] = 0
			i -= 2
			continue
		}
		stones[i-1] = stones[i] - stones[i-1]
		for j := i - 1; j >= 1; j-- {
			if stones[j] < stones[j-1] {
				stones[j], stones[j-1] = stones[j-1], stones[j]
			}
		}
		i -= 1
	}
	return stones[0]
}
