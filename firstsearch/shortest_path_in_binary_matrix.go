package firstsearch

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return -1
	}
	if grid[0][0] == 1 {
		return -1
	}
	if grid[n-1][n-1] == 1 {
		return -1
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	q := [][2]int{
		{0, 0},
	}
	cnt := 1
	dp[0][0] = 1
	illegal := func(x, y int) bool {
		if x < 0 {
			return true
		}
		if y < 0 {
			return true
		}
		if x >= n {
			return true
		}
		if y >= n {
			return true
		}
		if grid[x][y] == 1 {
			return true
		}
		if dp[x][y] > 0 {
			return true
		}
		return false
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for i := 0; i < len(tmp); i++ {
			cnt += 1
			x, y := tmp[i][0], tmp[i][1]
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					if illegal(x+j, y+k) {
						continue
					}
					dp[x+j][y+k] = cnt
					q = append(q, [2]int{x + j, y + k})
				}
			}
		}
	}
	return dp[n-1][n-1]
}
