package dp

import "math"

func minFallingPathSum(matrix [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := len(matrix)
	if n == 0 {
		return 0
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			val := matrix[i-1][j]
			if j > 0 {
				val = min(val, matrix[i-1][j-1])
			}
			if j < n-1 {
				val = min(val, matrix[i-1][j+1])
			}
			matrix[i][j] += val
		}
	}
	ans := math.MaxInt
	i := n - 1
	for j := 0; j < n; j++ {
		ans = min(ans, matrix[i][j])
	}
	return ans
}
