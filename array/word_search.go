package array

func exist(board [][]byte, word string) bool {
	status := make([][]bool, len(board))
	for i := 0; i < len(status); i++ {
		status[i] = make([]bool, len(board[0]))
	}
	m := len(board)
	n := len(board[0])
	var bk func(i, j, index int) bool
	bk = func(i, j, index int) bool {
		if status[i][j] {
			return false
		}
		if board[i][j] != word[index] {
			return false
		}
		status[i][j] = true
		if index == len(word)-1 {
			return true
		}
		if i+1 < m && bk(i+1, j, index+1) {
			return true
		}
		if i-1 >= 0 && bk(i-1, j, index+1) {
			return true
		}
		if j+1 < n && bk(i, j+1, index+1) {
			return true
		}
		if j-1 >= 0 && bk(i, j-1, index+1) {
			return true
		}
		status[i][j] = false
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if bk(i, j, 0) {
				return true
			}
		}
	}
	return false
}
