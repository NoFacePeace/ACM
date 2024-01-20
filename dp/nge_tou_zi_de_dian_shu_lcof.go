package dp

func dicesProbability(n int) []float64 {
	base := 1.0 / 6
	ans := []float64{base, base, base, base, base, base}
	i := 1
	for i < n {
		tmp := ans
		ans = nil
		left, right := 0, 0
		val := 0.0
		for left < len(tmp) {
			val += tmp[right] * base
			ans = append(ans, val)
			if right+1 < len(tmp) {
				right++
			} else {
				val -= tmp[left] * base
				val -= tmp[right] * base
				left++
			}
			if right-left >= 6 {
				val -= tmp[left] * base
				left++
			}
		}
		i++
	}
	return ans
}
