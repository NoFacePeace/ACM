package math

func nthMagicalNumber(n int, a int, b int) int {
	var mod int = 1e9 + 7
	l := min(a, b)
	r := n * min(a, b)
	c := a / gcd(a, b) * b
	for l < r {
		mid := (l + r) / 2
		cnt := mid/a + mid/b - mid/c
		if cnt >= n {
			r = mid
			continue
		}
		l = mid + 1
	}
	return r % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
