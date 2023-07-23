package firstsearch

func nearestExit(maze [][]byte, entrance []int) int {
	m := len(maze)
	if m == 0 {
		return 0
	}
	n := len(maze[0])
	if n == 0 {
		return 0
	}
	if maze[entrance[0]][entrance[1]] == '+' {
		return -1
	}
	q := [][]int{entrance}
	maze[entrance[0]][entrance[1]] = '+'
	dirs := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	ans := -1
	cnt := 0
	for len(q) > 0 {
		cnt++
		tmp := q
		q = nil
		for _, v := range tmp {
			i, j := v[0], v[1]
			for _, d := range dirs {
				x := i + d[0]
				y := j + d[1]
				if x < 0 || x >= m {
					continue
				}
				if y < 0 || y >= n {
					continue
				}
				if maze[x][y] == '+' {
					continue
				}
				if x == 0 || x == m-1 {
					ans = cnt
					break
				}
				if y == 0 || y == n-1 {
					ans = cnt
					break
				}
				maze[x][y] = '+'
				q = append(q, []int{x, y})
			}
			if ans != -1 {
				q = nil
				break
			}
		}
	}
	return ans
}
