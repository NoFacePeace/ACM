package dp

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	if n == 0 {
		return 1
	}
	if n == 1 {
		return cost[0]
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	if n == 2 {
		return min(cost[0], cost[1])
	}
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1], dp[i-2])
		if i != n {
			dp[i] += cost[i]
		}
	}
	return dp[n]
}
