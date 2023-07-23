package dp

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for i := v; i <= amount; i++ {
			dp[i] += dp[i-v]
		}
	}
	return dp[amount]
}
