package array

func largest1BorderedSquare(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	maxL := m
	if m > n {
		maxL = n
	}
	ans := 1
	isSquare := func(x, y, l int) bool {
		for i := y; i < y+l; i++ {
			if grid[x][i] != 1 {
				return false
			}
			if grid[x+l-1][i] != 1 {
				return false
			}
		}
		for i := x; i < x+l; i++ {
			if grid[i][y] != 1 {
				return false
			}
			if grid[i][y+l-1] != 1 {
				return false
			}
		}
		return true
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max + 1
			for ans <= maxL {
				if i+ans > m {
					break
				}
				if j+ans > n {
					break
				}
				if isSquare(i, j, ans) {
					max = ans
				}
				ans++
			}
		}
	}
	return max * max
}
