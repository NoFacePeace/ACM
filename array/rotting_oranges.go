package array

func orangesRotting(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return -1
	}
	n := len(grid[0])
	if n == 0 {
		return -1
	}
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}
	ans := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		time := 0
		for _, v := range tmp {
			i, j := v[0], v[1]
			if grid[i][j] == 0 {
				continue
			}
			grid[i][j] = 0
			time = 1
			if i-1 >= 0 && grid[i-1][j] == 1 {
				q = append(q, []int{i - 1, j})
			}
			if i+1 < m && grid[i+1][j] == 1 {
				q = append(q, []int{i + 1, j})
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				q = append(q, []int{i, j - 1})
			}
			if j+1 < n && grid[i][j+1] == 1 {
				q = append(q, []int{i, j + 1})
			}
		}
		ans += time
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	if ans == 0 {
		return 0
	}
	return ans - 1
}
