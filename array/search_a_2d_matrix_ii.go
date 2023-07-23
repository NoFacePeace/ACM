package array

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1
	for {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			if j > 0 {
				j--
				continue
			}
			return false
		}
		if i < m-1 {
			i++
			continue
		}
		return false
	}
}
