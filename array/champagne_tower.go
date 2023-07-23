package array

func champagneTower(poured int, query_row int, query_glass int) float64 {
	arr := [][]float64{}
	column := 1
	for i := 0; i <= query_row; i++ {
		arr = append(arr, []float64{})
		arr[i] = make([]float64, column)
		column++
		if i == 0 {
			arr[i][0] = float64(poured)
			continue
		}
		l := len(arr[i])
		for j := 0; j < l; j++ {
			left, right := 0.0, 0.0
			if j > 0 {
				left = (arr[i-1][j-1] - 1) / 2
			}
			if j < l-1 {
				right = (arr[i-1][j] - 1) / 2
			}
			if left > 0 {
				arr[i][j] += left
			}
			if right > 0 {
				arr[i][j] += right
			}
		}

	}
	if arr[query_row][query_glass] > 1 {
		return 1.0
	}
	return arr[query_row][query_glass]
}
