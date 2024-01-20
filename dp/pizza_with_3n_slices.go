package dp

func maxSizeSlices(slices []int) int {
	n := len(slices)
	k := n / 3
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	f := func(slices []int) int {
		n := len(slices)
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, k+1)
		}
		for i := 0; i < n; i++ {
			for j := 1; j < k+1; j++ {
				if i == 0 && j == 1 {
					dp[i][j] = slices[i]
					continue
				}
				if i == 1 && j == 1 {
					dp[i][j] = max(slices[0], slices[1])
					continue
				}
				if i+1 <= j {
					break
				}
				dp[i][j] = max(dp[i-2][j-1]+slices[i], dp[i-1][j])
			}
		}
		return dp[n-1][k]
	}
	return max(f(slices[1:]), f(slices[:n-1]))
}
