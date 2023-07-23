package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	top := 0
	bottom := m - 1
	for top < bottom {
		mid := (top + bottom) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target {
			bottom = mid - 1
			continue
		}
		if matrix[mid][n-1] == target {
			return true
		}
		if matrix[mid][n-1] > target {
			top = mid
			break
		}
		top = mid + 1
	}
	for i := 0; i < n; i++ {
		if matrix[top][i] == target {
			return true
		}
	}
	return false
}
