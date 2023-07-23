package array

import "math"

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			max := math.MinInt
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if grid[i+x][j+y] > max {
						max = grid[i+x][j+y]
					}
				}
			}
			ans[i-1] = append(ans[i-1], max)
		}
	}
	return ans
}
