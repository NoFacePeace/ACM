package firstsearch

func shortestBridge(grid [][]int) int {
	n := len(grid)
	x, y := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				x = i
				y = j
				break
			}
		}
		if x != -1 && y != -1 {
			break
		}
	}
	if x == -1 && y == -1 {
		return 0
	}
	q := [][]int{}
	arr := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var f func(i, j int)
	f = func(i, j int) {
		if i < 0 {
			return
		}
		if j < 0 {
			return
		}
		if i >= n {
			return
		}
		if j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		if grid[i][j] == 2 {
			return
		}
		q = append(q, []int{i, j})
		grid[i][j] = 2
		for _, v := range arr {
			f(i+v[0], j+v[1])
		}
	}
	f(x, y)
	cnt := -1
	ok := false
	for len(q) > 0 {
		cnt++
		tmp := q
		q = nil
		for _, v := range tmp {
			i, j := v[0], v[1]
			for _, w := range arr {
				if i+w[0] < 0 || i+w[0] >= n {
					continue
				}
				if j+w[1] < 0 || j+w[1] >= n {
					continue
				}
				if grid[i+w[0]][j+w[1]] == 2 {
					continue
				}
				if grid[i+w[0]][j+w[1]] == 1 {
					q = nil
					ok = true
					break
				}
				grid[i+w[0]][j+w[1]] = 2
				q = append(q, []int{i + w[0], j + w[1]})
			}
			if ok {
				break
			}
		}
	}
	return cnt
}
