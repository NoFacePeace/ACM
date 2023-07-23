package str

func printBin(num float64) string {
	var f func(num, bit float64)
	ans := "0."
	f = func(num, bit float64) {
		if num == bit {
			ans += "1"
			return
		}
		if num < bit {
			ans += "0"
			f(num, bit/2)
			return
		}
		dist := num - bit
		ans += "1"
		f(dist, bit/2)
	}
	f(num, 1.0/2.0)
	if len(ans) <= 34 {
		return ans
	}
	return "ERROR"
}
