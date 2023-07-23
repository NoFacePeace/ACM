package dp

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-v]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
