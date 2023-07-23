package array

func calculateTax(brackets [][]int, income int) float64 {
	ans := 0
	upper := 0
	for _, v := range brackets {
		if income <= v[0] {
			ans += (income - upper) * v[1]
			break
		}
		ans += (v[0] - upper) * v[1]
		upper = v[0]
	}
	return float64(ans) / 100.0
}
