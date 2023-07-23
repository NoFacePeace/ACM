package dp

import "math"

func minimumTotal(triangle [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < len(triangle); i++ {
		if i == 0 {
			continue
		}
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
				continue
			}
			if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
				continue
			}
			val := triangle[i-1][j]
			val = min(val, triangle[i-1][j-1])
			triangle[i][j] += val
		}
	}
	i := len(triangle) - 1
	ans := math.MaxInt
	for j := 0; j < len(triangle[i]); j++ {
		ans = min(ans, triangle[i][j])
	}
	return ans
}
