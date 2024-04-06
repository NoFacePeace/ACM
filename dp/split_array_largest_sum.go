package dp

func splitArray(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			sum[i] = nums[i]
			continue
		}
		sum[i] += sum[i-1] + nums[i]
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(end, k int) int
	dfs = func(end, k int) int {
		if k == 1 {
			dp[end][k] = sum[end]
			return sum[end]
		}
		if dp[end][k] != -1 {
			return dp[end][k]
		}
		val := -1
		for i := end; i >= 0; i-- {
			if i+1 < k {
				break
			}
			mx := max(sum[end]-sum[i]+nums[i], dfs(i-1, k-1))
			if val == -1 {
				val = mx
				continue
			}
			if mx < val {
				val = mx
			}
		}
		dp[end][k] = val
		return val
	}
	dfs(n-1, k)
	return dp[n-1][k]
}
