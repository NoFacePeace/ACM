package dp

func maxProfitFee(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy := -prices[0]
	sell := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		buy, sell = max(buy, sell-prices[i]), max(sell, buy+prices[i]-fee)
	}
	return sell
}
