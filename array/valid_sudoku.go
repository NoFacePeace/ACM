package array

func isValidSudoku(board [][]byte) bool {
	check := func(arr []byte) bool {
		m := map[byte]bool{}
		for i := range arr {
			if arr[i] == '.' {
				continue
			}
			if _, ok := m[arr[i]]; ok {
				return false
			}
			m[arr[i]] = true
		}
		return true
	}
	for i := range board {
		if !check(board[i]) {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		arr := []byte{}
		for j := 0; j < 9; j++ {
			arr = append(arr, board[j][i])
		}
		if !check(arr) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			arr := []byte{}
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					arr = append(arr, board[x][y])
				}
			}
			if !check(arr) {
				return false
			}
		}
	}
	return true
}
