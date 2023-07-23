package array

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	if m == 0 {
		return 0
	}
	n := len(grid1[0])
	if n == 0 {
		return 0
	}
	var f func(i, j, num int)
	f = func(i, j, num int) {
		if i < 0 {
			return
		}
		if j < 0 {
			return
		}
		if i >= m {
			return
		}
		if j >= n {
			return
		}
		if grid1[i][j] == 0 {
			return
		}
		if grid1[i][j] > 1 {
			return
		}
		grid1[i][j] = num
		f(i+1, j, num)
		f(i-1, j, num)
		f(i, j-1, num)
		f(i, j+1, num)
	}
	num := 2
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == 0 {
				continue
			}
			if grid1[i][j] > 1 {
				continue
			}
			f(i, j, num)
			num += 1
		}
	}
	ans := 0
	var f1 func(i, j, num int) bool
	f1 = func(i, j, num int) bool {
		if i < 0 {
			return true
		}
		if j < 0 {
			return true
		}
		if i >= m {
			return true
		}
		if j >= n {
			return true
		}
		if grid2[i][j] == 0 {
			return true
		}
		grid2[i][j] = 0
		r := true
		if grid1[i][j] != num {
			r = false
		}
		if !f1(i+1, j, num) {
			r = false
		}
		if !f1(i-1, j, num) {
			r = false
		}
		if !f1(i, j-1, num) {
			r = false
		}
		if !f1(i, j+1, num) {
			r = false
		}
		return r
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 0 {
				continue
			}
			if grid1[i][j] == 0 {
				continue
			}
			if f1(i, j, grid1[i][j]) {
				ans += 1
			}
		}
	}
	return ans
}
