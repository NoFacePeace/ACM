package dp

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}
		if i == 1 {
			dp[i] = max(dp[i-1], nums[i])
			continue
		}
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}
