package array

func rotate(matrix [][]int) {
	level := 0
	l := len(matrix) - 1
	for level < len(matrix)/2 {
		for k := level; k < l-level; k++ {
			x := level
			y := k
			tmp := matrix[x][y]
			for {
				i, j := find(x, y, l)
				matrix[i][j], tmp = tmp, matrix[i][j]
				if i == level && j == k {
					break
				}
				x = i
				y = j
			}
		}
		level++
	}
}

func find(x, y, l int) (int, int) {
	i := y
	j := l - x
	return i, j
}
