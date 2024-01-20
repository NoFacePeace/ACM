package backtrack

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	var f func(x, y, p int) bool
	f = func(x, y, p int) bool {
		if p >= len(word) {
			return true
		}
		if x < 0 || x >= m {
			return false
		}
		if y < 0 || y >= n {
			return false
		}
		if visit[x][y] {
			return false
		}
		if board[x][y] != word[p] {
			return false
		}
		visit[x][y] = true
		if f(x+1, y, p+1) {
			return true
		}
		if f(x-1, y, p+1) {
			return true
		}
		if f(x, y+1, p+1) {
			return true
		}
		if f(x, y-1, p+1) {
			return true
		}
		visit[x][y] = false
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f(i, j, 0) {
				return true
			}
		}
	}
	return false
}
