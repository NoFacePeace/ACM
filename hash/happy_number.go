package hash

func isHappy(n int) bool {
	m := map[int]bool{}
	f := func(n int) int {
		num := 0
		for n != 0 {
			num += (n % 10) * (n % 10)
			n /= 10
		}
		return num
	}
	for n != 1 {
		m[n] = true
		n = f(n)
		if m[n] {
			return false
		}
	}
	return true
}
