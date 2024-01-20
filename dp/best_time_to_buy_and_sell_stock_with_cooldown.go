package dp

func maxProfitCoolDown(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	on := -prices[0]
	off := 0
	cold := 0
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	for i := 1; i < n; i++ {
		on, off, cold = max(on, off-prices[i]), max(off, cold), on+prices[i]
	}
	return max(off, cold)
}
