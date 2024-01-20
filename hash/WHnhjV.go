package hash

import "math"

func giveGem(gem []int, operations [][]int) int {
	m := map[int]int{}
	for k, v := range gem {
		m[k] = v
	}
	for _, op := range operations {
		g1 := op[0]
		g2 := op[1]
		val := m[g1] / 2
		m[g2] += val
		m[g1] -= val
	}
	min := math.MaxInt
	max := math.MinInt
	for _, v := range m {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}
