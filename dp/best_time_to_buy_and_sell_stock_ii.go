package dp

func maxProfitII(prices []int) int {
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
	for _, v := range prices {
		buy, sell = max(buy, sell-v), max(sell, buy+v)
	}
	return sell
}
