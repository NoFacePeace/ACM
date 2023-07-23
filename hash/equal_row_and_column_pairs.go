package hash

import "strconv"

func equalPairs(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := map[string]int{}
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < n; j++ {
			ss := strconv.Itoa(grid[i][j])
			s += ss + "#"
		}
		m[s]++
	}
	ans := 0
	for i := 0; i < n; i++ {
		s := ""
		for j := 0; j < n; j++ {
			ss := strconv.Itoa(grid[j][i])
			s += ss + "#"
		}
		ans += m[s]
	}
	return ans
}
