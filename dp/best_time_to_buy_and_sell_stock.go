package dp

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min := prices[0]
	max := 0
	for k := range prices {
		if prices[k] < min {
			min = prices[k]
			continue
		}
		if prices[k]-min > max {
			max = prices[k] - min
		}
	}
	return max
}
