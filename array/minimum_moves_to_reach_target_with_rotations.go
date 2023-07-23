package array

func minimumMoves(grid [][]int) int {
	n := len(grid)
	dist := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([][2]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = [2]int{-1, -1}
		}
	}
	dist[0][0][0] = 0
	type pair struct {
		x      int
		y      int
		status int
	}
	q := []pair{
		{0, 0, 0},
	}
	right := func(v pair) bool {
		if v.status == 0 {
			if v.y+2 >= n {
				return false
			}
			if grid[v.x][v.y+2] == 1 {
				return false
			}
			if dist[v.x][v.y+1][0] != -1 {
				return false
			}
			return true
		}
		if v.y+1 >= n {
			return false
		}
		if grid[v.x][v.y+1] == 1 {
			return false
		}
		if grid[v.x+1][v.y+1] == 1 {
			return false
		}
		if dist[v.x][v.y+1][1] != -1 {
			return false
		}
		return true
	}
	bottom := func(v pair) bool {
		if v.status == 0 {
			if v.x+1 >= n {
				return false
			}
			if grid[v.x+1][v.y] == 1 {
				return false
			}
			if grid[v.x+1][v.y+1] == 1 {
				return false
			}
			if dist[v.x+1][v.y][0] != -1 {
				return false
			}
			return true
		}
		if v.x+2 >= n {
			return false
		}
		if grid[v.x+2][v.y] == 1 {
			return false
		}
		if dist[v.x+1][v.y][1] != -1 {
			return false
		}
		return true
	}
	wise := func(v pair) bool {
		if v.status != 0 {
			return false
		}
		if v.x+1 >= n {
			return false
		}
		if grid[v.x+1][v.y] == 1 {
			return false
		}
		if grid[v.x+1][v.y+1] == 1 {
			return false
		}
		if dist[v.x][v.y][1] != -1 {
			return false
		}
		return true
	}
	antiwise := func(v pair) bool {
		if v.status != 1 {
			return false
		}
		if v.y+1 >= n {
			return false
		}
		if grid[v.x][v.y+1] == 1 {
			return false
		}
		if grid[v.x+1][v.y+1] == 1 {
			return false
		}
		if dist[v.x][v.y][0] != -1 {
			return false
		}
		return true
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if right(v) {
				q = append(q, pair{v.x, v.y + 1, v.status})
				dist[v.x][v.y+1][v.status] = dist[v.x][v.y][v.status] + 1
			}
			if bottom(v) {
				q = append(q, pair{v.x + 1, v.y, v.status})
				dist[v.x+1][v.y][v.status] = dist[v.x][v.y][v.status] + 1
			}
			if wise(v) {
				q = append(q, pair{v.x, v.y, 1})
				dist[v.x][v.y][1] = dist[v.x][v.y][0] + 1
			}
			if antiwise(v) {
				q = append(q, pair{v.x, v.y, 0})
				dist[v.x][v.y][0] = dist[v.x][v.y][1] + 1
			}
		}
	}
	return dist[n-1][n-2][0]
}
