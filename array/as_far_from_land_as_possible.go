package array

func maxDistance(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	q := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				q = append(q, []int{i, j})
			}
		}
	}
	cnt := 1
	ans := -1
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			x, y := v[0], v[1]
			if grid[x][y] > 0 {
				continue
			}
			grid[x][y] = cnt
			ans = cnt
			if x-1 >= 0 && grid[x-1][y] == 0 {
				q = append(q, []int{x - 1, y})
			}
			if x+1 < n && grid[x+1][y] == 0 {
				q = append(q, []int{x + 1, y})
			}
			if y-1 >= 0 && grid[x][y-1] == 0 {
				q = append(q, []int{x, y - 1})
			}
			if y+1 < n && grid[x][y+1] == 0 {
				q = append(q, []int{x, y + 1})
			}
		}
		cnt++
	}
	if ans == 1 {
		return -1
	}
	return ans - 1
}
