package hash

import "math"

func flipgame(fronts []int, backs []int) int {
	m := map[int]bool{}
	min := math.MaxInt
	n := len(fronts)
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			m[fronts[i]] = true
		}
	}
	for i := 0; i < n; i++ {
		f := fronts[i]
		if !m[f] && f < min {
			min = f
		}
		b := backs[i]
		if !m[b] && b < min {
			min = b

		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}
