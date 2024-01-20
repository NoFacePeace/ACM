package array

func buyChoco(prices []int, money int) int {
	n := len(prices)
	if n < 2 {
		return money
	}
	mn := prices[0]
	mx := prices[1]
	if mx < mn {
		mx, mn = mn, mx
	}
	for i := 2; i < n; i++ {
		if prices[i] > mx {
			continue
		}
		mx = prices[i]
		if mx < mn {
			mx, mn = mn, mx
		}
	}
	if mx+mn > money {
		return money
	}
	return money - mx - mn
}
