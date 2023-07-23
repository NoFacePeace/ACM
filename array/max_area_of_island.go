package array

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	max := 0
	var f func(i, j int) int
	f = func(i, j int) int {
		if i < 0 {
			return 0
		}
		if i >= m {
			return 0
		}
		if j < 0 {
			return 0
		}
		if j >= n {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + f(i+1, j) + f(i-1, j) + f(i, j+1) + f(i, j-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			val := f(i, j)
			if val > max {
				max = val
			}
		}
	}
	return max
}
