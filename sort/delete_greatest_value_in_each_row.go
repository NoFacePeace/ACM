package sort

import "sort"

func deleteGreatestValue(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		sort.Ints(grid[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		max := 0
		for j := 0; j < m; j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		ans += max
	}
	return ans
}
