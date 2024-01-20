package array

func diagonalSum(mat [][]int) int {
	n := len(mat)
	x1 := 0
	y1 := 0
	x2 := 0
	y2 := n - 1
	sum := 0
	for x1 < n {
		if x1 == x2 && y1 == y2 {
			sum += mat[x1][y1]
		} else {
			sum += mat[x1][y1]
			sum += mat[x2][y2]
		}
		x1++
		y1++
		x2++
		y2--
	}
	return sum
}
