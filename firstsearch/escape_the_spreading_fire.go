package firstsearch

func maximumMinutes(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				grid[i][j] = -1
			}
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			x, y := v[0], v[1]
			if x-1 >= 0 && grid[x-1][y] == 0 {
				grid[x-1][y] = grid[x][y] + 1
				q = append(q, []int{x - 1, y})
			}
			if x+1 < m && grid[x+1][y] == 0 {
				grid[x+1][y] = grid[x][y] + 1
				q = append(q, []int{x + 1, y})
			}
			if y-1 >= 0 && grid[x][y-1] == 0 {
				grid[x][y-1] = grid[x][y] + 1
				q = append(q, []int{x, y - 1})
			}
			if y+1 < n && grid[x][y+1] == 0 {
				grid[x][y+1] = grid[x][y] + 1
				q = append(q, []int{x, y + 1})
			}
		}
	}
	var bfs func(x, y, cnt int) (bool, int)
	bfs = func(x, y, cnt int) (bool, int) {
		q := [][]int{{x, y}}
		visited := make([][]bool, m)
		for i := 0; i < m; i++ {
			visited[i] = make([]bool, n)
		}
		visited[x][y] = true
		pairs := [][]int{
			{1, 0},
			{-1, 0},
			{0, 1},
			{0, -1},
		}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, p := range pairs {
					i, j := v[0]+p[0], v[1]+p[1]
					if i < 0 || i >= m {
						continue
					}
					if j < 0 || j >= n {
						continue
					}
					if visited[i][j] {
						continue
					}
					if i == m-1 && j == n-1 {
						if grid[i][j] == 0 {
							return true, 0
						}
						if cnt <= grid[i][j] {
							return true, grid[i][j] - cnt
						}
					}
					if grid[i][j] != 0 && cnt >= grid[i][j] {
						continue
					}
					q = append(q, []int{i, j})
					visited[i][j] = true
				}
			}
			cnt++
		}
		return false, -1
	}
	ok, cnt := bfs(0, 0, 2)
	if !ok {
		return -1
	}
	if grid[m-1][n-1] == 0 {
		return int(1e9)
	}
	ok, _ = bfs(0, 0, cnt+2)
	if ok {
		return cnt
	}
	return cnt - 1
}
