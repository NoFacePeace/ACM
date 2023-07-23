package dp

func matrixBlockSum(mat [][]int, k int) [][]int {
	ans := [][]int{}
	m := len(mat)
	if m == 0 {
		return ans
	}
	n := len(mat[0])
	if n == 0 {
		return ans
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				mat[i][j] += mat[i][j-1]
				continue
			}
			if j == 0 {
				mat[i][j] += mat[i-1][j]
				continue
			}
			mat[i][j] += mat[i-1][j] + mat[i][j-1]
			mat[i][j] -= mat[i-1][j-1]
		}
	}
	ans = make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x1 := i + k
			if x1 >= m {
				x1 = m - 1
			}
			y1 := j + k
			if y1 >= n {
				y1 = n - 1
			}
			ans[i][j] = mat[x1][y1]
			x2 := i - k - 1
			y2 := j - k - 1
			if x2 >= 0 && y2 >= 0 {
				ans[i][j] += mat[x2][y2]
			}
			if x2 >= 0 {
				ans[i][j] -= mat[x2][y1]
			}
			if y2 >= 0 {
				ans[i][j] -= mat[x1][y2]
			}
		}
	}
	return ans
}
