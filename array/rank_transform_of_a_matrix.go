package array

import "sort"

func matrixRankTransform(matrix [][]int) [][]int {
	type pair struct {
		val int
		x   int
		y   int
	}
	arr := []pair{}
	m := len(matrix)
	if m <= 0 {
		return matrix
	}
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr = append(arr, pair{matrix[i][j], i, j})
		}
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a].val < arr[b].val
	})
	row := make([]int, m)
	maxRow := make([]int, m)
	column := make([]int, n)
	maxColumn := make([]int, n)
	for _, v := range arr {
		var p int
		if row[v.x] > column[v.y] {
			if maxRow[v.x] == v.val {
				p = row[v.x]
			} else {
				p = row[v.x] + 1
			}
		} else {
			if maxColumn[v.y] == v.val {
				p = column[v.y]
			} else {
				p = column[v.y] + 1
			}
		}
		row[v.x] = p
		maxRow[v.x] = v.val
		column[v.y] = p
		maxColumn[v.y] = v.val
		matrix[v.x][v.y] = p
	}
	return matrix
}
