package array

func closedIsland(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	var f func(i, j int)
	f = func(i, j int) {
		if i < 0 {
			return
		}
		if j < 0 {
			return
		}
		if i >= m {
			return
		}
		if j >= n {
			return
		}
		if grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
		f(i+1, j)
		f(i-1, j)
		f(i, j+1)
		f(i, j-1)
	}
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			f(i, 0)
		}
		if grid[i][n-1] == 0 {
			f(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			f(0, j)
		}
		if grid[m-1][j] == 0 {
			f(m-1, j)
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				continue
			}
			ans++
			f(i, j)
		}
	}
	return ans
}
