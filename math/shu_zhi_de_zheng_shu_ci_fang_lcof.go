package math

func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var f func(n int) float64
	f = func(n int) float64 {
		if n == 1 {
			return x
		}
		val := f(n / 2)
		val *= val
		if n%2 != 0 {
			val *= x
		}
		return val
	}
	sign := true
	if n < 0 {
		n = -n
		sign = false
	}
	if sign {
		return f(n)
	}
	return 1 / f(n)
}
