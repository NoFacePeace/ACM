package array

func maxProfit1(prices []int) int {
	l := len(prices)
	dp := make([][3]int, l)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][2]-prices[i])
		dp[i][2] = dp[i-1][0]
	}
	return max(dp[l-1][0], dp[l-1][2])
}
