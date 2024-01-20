package greedy

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	min := prices[0]
	ans := 0
	for i := 1; i < n; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > ans {
			ans = prices[i] - min
		}
	}
	return ans
}
