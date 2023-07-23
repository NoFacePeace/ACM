package firstsearch

func closedIsland(grid [][]int) int {
	n := len(grid)
	if n <= 2 {
		return 0
	}
	m := len(grid[0])
	if m <= 2 {
		return 0
	}
	var fill func(x, y int)
	fill = func(x, y int) {
		if x < 0 {
			return
		}
		if y < 0 {
			return
		}
		if x >= n {
			return
		}
		if y >= m {
			return
		}
		if grid[x][y] == 1 {
			return
		}
		grid[x][y] = 1
		fill(x+1, y)
		fill(x-1, y)
		fill(x, y+1)
		fill(x, y-1)
	}
	for i := 0; i < n; i++ {
		if grid[i][0] == 0 {
			fill(i, 0)
		}
		if grid[i][m-1] == 0 {
			fill(i, m-1)
		}
	}
	for j := 0; j < m; j++ {
		if grid[0][j] == 0 {
			fill(0, j)
		}
		if grid[n-1][j] == 0 {
			fill(n-1, j)
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				fill(i, j)
				ans++
			}
		}
	}
	return ans
}
