package array

type pair struct {
	x int
	y int
}

func shortestBridge(grid [][]int) int {

	dirs := []pair{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	n := len(grid)
	q := []pair{}
	isLand := []pair{}
	final := false
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			grid[i][j] = -1
			q = append(q, pair{i, j})
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				isLand = append(isLand, p)
				for _, v := range dirs {
					x, y := p.x+v.x, p.y+v.y
					if x < 0 || x >= n || y < 0 || y >= n {
						continue
					}
					if grid[x][y] != 1 {
						continue
					}
					q = append(q, pair{x, y})
					grid[x][y] = -1
				}
			}
			final = true
			break
		}
		if final {
			break
		}
	}
	q = isLand
	step := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			for _, v := range dirs {
				x, y := p.x+v.x, p.y+v.y
				if x < 0 || x >= n || y < 0 || y >= n {
					continue
				}
				if grid[x][y] == 1 {
					return step
				}
				if grid[x][y] == 0 {
					grid[x][y] = -1
					q = append(q, pair{x, y})
				}
			}
		}
		step++
	}
	return step
}
