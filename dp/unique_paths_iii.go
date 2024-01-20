package dp

func uniquePathsIII(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	x, y := 0, 0
	zero := 0
	start := func() {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if grid[i][j] == 1 {
					x, y = i, j
				}
				if grid[i][j] == 0 {
					zero++
				}
			}
		}
	}
	start()
	ans := 0
	pair := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var dfs func(x, y, cnt int)
	dfs = func(x, y, cnt int) {
		if x < 0 {
			return
		}
		if x >= n {
			return
		}
		if y < 0 {
			return
		}
		if y >= m {
			return
		}
		if grid[x][y] == 2 {
			if cnt == zero+1 {
				ans++
			}
			return
		}
		if grid[x][y] == -1 {
			return
		}
		grid[x][y] = -1
		for _, p := range pair {
			dfs(x+p[0], y+p[1], cnt+1)
		}
		grid[x][y] = 0
	}
	dfs(x, y, 0)
	return ans
}
