package dp

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := nums[0]
	min := 0
	ans := max
	maxf := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	minf := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1; i < n; i++ {
		val := nums[i]
		mx := maxf(val-min, val)
		mn := minf(min, val-max)
		mn = minf(mn, min)
		mx = maxf(mx, max)
		max = mx
		min = mn
		ans = maxf(max, ans)
	}
	return int64(ans)
}
