package firstsearch

import "math"

func minPushBox(grid [][]byte) int {
	arr := [][2]int{}
	m := len(grid)
	n := len(grid[0])
	var px, py, bx, by int
	ans := math.MaxInt
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' {
				px = i
				py = j
			}
			if grid[i][j] == 'B' {
				bx = i
				by = j
			}
		}
	}
	overflow := func(x, y int) bool {
		if x < 0 {
			return true
		}
		if y < 0 {
			return true
		}
		if x >= m {
			return true
		}
		if y >= n {
			return true
		}
		return false
	}
	cover := func() {
		var dfs func(x, y int)
		dfs = func(x, y int) {
			if overflow(x, y) {
				return
			}
			if grid[x][y] == '.' {
				grid[x][y] = 'S'
				dfs(x, y+1)
				dfs(x, y-1)
				dfs(x+1, y)
				dfs(x-1, y)
			}
		}
		grid[px][py] = '.'
		dfs(px, py)
	}
	clear := func() {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 'S' {
					grid[i][j] = '.'
				}
			}
		}
		grid[px][py] = 'S'
	}
	check := func() bool {
		cover()
		defer clear()
		for _, v := range arr {
			if grid[v[0]][v[1]] == 'B' || grid[v[0]][v[1]] == 'S' {
				continue
			}
			return false
		}
		return true
	}
	push := func() {
		var dfs func(x, y, cnt int)
		dfs = func(x, y, cnt int) {
			if overflow(x, y) {
				return
			}
			if grid[x][y] == 'T' {
				if cnt < ans {
					if check() {
						ans = cnt
					}
				}
				return
			}
			if grid[x][y] == '.' {
				grid[x][y] = 'B'
				arr = append(arr, [2]int{x, y - 1})
				dfs(x, y+1, cnt+1)
				arr = arr[:len(arr)-1]
				arr = append(arr, [2]int{x, y + 1})
				dfs(x, y-1, cnt+1)
				arr = arr[:len(arr)-1]
				arr = append(arr, [2]int{x - 1, y})
				dfs(x+1, y, cnt+1)
				arr = arr[:len(arr)-1]
				arr = append(arr, [2]int{x + 1, y})
				dfs(x-1, y, cnt+1)
				arr = arr[:len(arr)-1]
				grid[x][y] = '.'
			}
		}
		grid[bx][by] = '.'
		dfs(bx, by, 0)
	}
	push()
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
