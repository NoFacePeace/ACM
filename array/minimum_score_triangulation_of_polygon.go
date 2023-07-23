package array

import "math"

func minScoreTriangulation(values []int) int {
	m := map[int]int{}
	l := len(values)
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i+2 > j {
			return 0
		}
		if i+2 == j {
			return values[i] * values[i+1] * values[i+2]
		}
		if m[i*l+j] != 0 {
			return m[i*l+j]
		}
		min := math.MaxInt
		for k := i + 1; k < j; k++ {
			val := values[i]*values[k]*values[j] + dp(i, k) + dp(k, j)
			if val < min {
				min = val
			}
		}
		m[i*l+j] = min
		return min
	}
	return dp(0, l-1)
}
