package array

func maxProfit2(prices []int, fee int) int {
	l := len(prices)
	dp := make([][2]int, l)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[l-1][0]
}
