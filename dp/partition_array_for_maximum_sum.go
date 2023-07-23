package dp

func maxSumAfterPartitioning(arr []int, k int) int {
	l := len(arr)
	dp := make([]int, l+1)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < l; i++ {
		dp[i+1] = dp[i] + arr[i]
		mx := arr[i]
		for j := 1; j < k && i-j >= 0; j++ {
			mx = max(mx, arr[i-j])
			dp[i+1] = max(dp[i+1], dp[i-j]+mx*(j+1))
		}
	}
	return dp[l]
}
