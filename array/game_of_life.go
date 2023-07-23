package array

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	has := func(x, y int) int {
		count := 0
		for i := x - 1; i < x+2; i++ {
			if i < 0 || i >= m {
				continue
			}
			for j := y - 1; j < y+2; j++ {
				if j < 0 || j >= n {
					continue
				}
				if i == x && j == y {
					continue
				}
				if board[i][j] >= 1 {
					count++
				}
			}
		}
		return count
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := has(i, j)
			if board[i][j] == 0 {
				board[i][j] = -count
				continue
			}
			board[i][j] = count + 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 0 {
				continue
			}
			if board[i][j] < 0 {
				if board[i][j] == -3 {
					board[i][j] = 1
					continue
				}
				board[i][j] = 0
				continue
			}
			if board[i][j] == 3 || board[i][j] == 4 {
				board[i][j] = 1
				continue
			}
			board[i][j] = 0
		}
	}
}
