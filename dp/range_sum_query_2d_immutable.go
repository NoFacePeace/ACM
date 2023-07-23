package dp

type NumMatrix struct {
	dp [][]int
	m  int
	n  int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	if m == 0 {
		return NumMatrix{}
	}
	n := len(matrix[0])
	if n == 0 {
		return NumMatrix{}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				matrix[i][j] += matrix[i][j-1]
				continue
			}
			if j == 0 {
				matrix[i][j] += matrix[i-1][j]
				continue
			}
			matrix[i][j] += matrix[i-1][j] + matrix[i][j-1] - matrix[i-1][j-1]
		}
	}
	return NumMatrix{
		m:  m,
		n:  n,
		dp: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := this.dp[row2][col2]
	if row1-1 >= 0 && col1-1 >= 0 {
		ans += this.dp[row1-1][col1-1]
	}
	if row1-1 >= 0 {
		ans -= this.dp[row1-1][col2]
	}
	if col1-1 >= 0 {
		ans -= this.dp[row2][col1-1]
	}
	return ans
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
