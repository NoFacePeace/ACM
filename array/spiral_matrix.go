package array

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	var f func(x, y, m, n int)
	f = func(x, y, m, n int) {
		if x > m || y > n {
			return
		}
		for i := y; i <= n; i++ {
			ans = append(ans, matrix[x][i])
		}
		for i := x + 1; i <= m; i++ {
			ans = append(ans, matrix[i][n])
		}
		for i := n - 1; i >= y && x != m; i-- {
			ans = append(ans, matrix[m][i])
		}
		for i := m - 1; i > x && y != n; i-- {
			ans = append(ans, matrix[i][y])
		}
		f(x+1, y+1, m-1, n-1)
	}
	f(0, 0, len(matrix)-1, len(matrix[0])-1)
	return ans
}
