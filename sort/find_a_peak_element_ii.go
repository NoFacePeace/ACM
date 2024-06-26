package sort

func findPeakGrid(mat [][]int) []int {
	maxElement := func(row []int) int {
		i := 0
		for j := range row {
			if row[i] < row[j] {
				i = j
			}
		}
		return i
	}
	m := len(mat)
	low, high := 0, m-1
	for low <= high {
		i := (low + high) / 2
		j := maxElement(mat[i])
		if i-1 >= 0 && mat[i][j] < mat[i-1][j] {
			high = i - 1
			continue
		}
		if i+1 < m && mat[i][j] < mat[i+1][j] {
			low = i + 1
			continue
		}
		return []int{i, j}
	}
	return nil
}
