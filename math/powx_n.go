package math

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	status := false
	if n < 0 {
		n = -n
		status = true
	}
	tmp := 1.0
	for n != 1 {
		if n&1 != 0 {
			tmp *= x
			n--
			continue
		}
		x *= x
		n = n / 2
	}
	if status {
		return 1 / (x * tmp)
	}
	return x * tmp
}
