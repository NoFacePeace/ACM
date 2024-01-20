package dp

func maxProfitIII(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy1 := -prices[0]
	sell1 := 0
	buy2 := -prices[0]
	sell2 := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, v := range prices {
		buy1, sell1, buy2, sell2 = max(-v, buy1), max(sell1, buy1+v), max(buy2, sell1-v), max(sell2, buy2+v)
	}
	return max(sell1, sell2)
}
