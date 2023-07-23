package array

func numIslands(grid [][]byte) int {
	count := 0
	m := len(grid)
	n := len(grid[0])
	var bk func(x, y int)
	bk = func(x, y int) {
		if grid[x][y] == '2' {
			return
		}
		if grid[x][y] == '0' {
			return
		}
		grid[x][y] = '2'
		if x-1 >= 0 {
			bk(x-1, y)
		}
		if x+1 < m {
			bk(x+1, y)
		}
		if y-1 >= 0 {
			bk(x, y-1)
		}
		if y+1 < n {
			bk(x, y+1)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '2' || grid[i][j] == '0' {
				continue
			}
			count++
			bk(i, j)
		}
	}
	return count
}
