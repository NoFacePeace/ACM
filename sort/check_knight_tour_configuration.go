package sort

import "sort"

func checkValidGrid(grid [][]int) bool {
	type cell struct {
		x    int
		y    int
		time int
	}
	cells := []cell{}
	n := len(grid)
	if n == 0 {
		return true
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c := cell{
				x:    i,
				y:    j,
				time: grid[i][j],
			}
			cells = append(cells, c)
		}
	}
	sort.Slice(cells, func(a, b int) bool {
		return cells[a].time < cells[b].time
	})
	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}
	x, y := cells[0].x, cells[0].y
	if x != 0 || y != 0 {
		return false
	}
	for i := 1; i < len(cells); i++ {
		x1, y1 := cells[i].x, cells[i].y
		if abs(x, x1) == 2 && abs(y, y1) == 1 {
			x, y = x1, y1
			continue
		}
		if abs(x, x1) == 1 && abs(y, y1) == 2 {
			x, y = x1, y1
			continue
		}
		return false
	}
	return true
}
