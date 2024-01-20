package sort

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	i := n - 1
	j := 0
	for i >= 0 && j < m {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}
