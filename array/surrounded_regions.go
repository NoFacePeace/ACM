package array

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	status := make([][]bool, m)
	for i := 0; i < m; i++ {
		status[i] = make([]bool, n)
	}
	var bk func(i, j int)
	bk = func(i, j int) {
		if status[i][j] {
			return
		}
		if board[i][j] != 'O' {
			return
		}
		status[i][j] = true
		if i+1 < m {
			bk(i+1, j)
		}
		if i-1 >= 0 {
			bk(i-1, j)
		}
		if j+1 < n {
			bk(i, j+1)
		}
		if j-1 >= 0 {
			bk(i, j-1)
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			bk(i, 0)
		}

		if board[i][n-1] == 'O' {
			bk(i, n-1)
		}

	}
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			bk(0, i)
		}

		if board[m-1][i] == 'O' {
			bk(m-1, i)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if status[i][j] {
				continue
			}
			board[i][j] = 'X'
		}
	}
}
