package dp

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	top := make([][]int, m)
	left := make([][]int, m)
	dp := make([][]int, m)
	max := 0
	for i := 0; i < len(matrix); i++ {
		top[i] = make([]int, n)
		left[i] = make([]int, n)
		dp[i] = make([]int, n)
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 && j == 0 {
				if matrix[i][j] == '1' {
					left[i][j] = 1
					top[i][j] = 1
					dp[i][j] = 1
					if max == 0 {
						max = 1
					}
				}
				continue
			}
			if i == 0 {
				if matrix[i][j] == '1' {
					left[i][j] = left[i][j-1] + 1
					top[i][j] = 1
					dp[i][j] = 1
					if max == 0 {
						max = 1
					}
				}
				continue
			}
			if j == 0 {
				if matrix[i][j] == '1' {
					top[i][j] = top[i-1][j] + 1
					dp[i][j] = 1
					left[i][j] = 1
					if max == 0 {
						max = 1
					}
				}
				continue
			}
			if matrix[i][j] == '0' {
				continue
			}
			left[i][j] = left[i][j-1] + 1
			top[i][j] = top[i-1][j] + 1
			min := left[i][j]
			if min > top[i][j] {
				min = top[i][j]
			}
			if min > dp[i-1][j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = min
			}
			if dp[i][j] > max {
				max = dp[i][j]
			}
		}
	}
	return max * max
}
