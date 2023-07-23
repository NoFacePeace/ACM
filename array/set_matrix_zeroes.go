package array

func setZeroes(matrix [][]int) {
	iarr := []int{}
	jarr := []int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				iarr = append(iarr, i)
				jarr = append(jarr, j)
			}
		}
	}
	for i := 0; i < len(iarr); i++ {
		for j := 0; j < len(matrix[iarr[i]]); j++ {
			matrix[iarr[i]][j] = 0
		}
	}
	for i := 0; i < len(jarr); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[j][jarr[i]] = 0
		}
	}
}
