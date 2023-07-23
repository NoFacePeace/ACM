package array

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	x := 0
	for x < n {
		if grid[x][x] == 0 {
			return false
		}
		grid[x][x] = -1
		x++
	}
	x = 0
	for x < n {
		if grid[x][n-x-1] == 0 {
			return false
		}
		grid[x][n-x-1] = -1
		x++
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 && grid[i][j] != -1 {
				return false
			}
		}
	}
	return true
}
