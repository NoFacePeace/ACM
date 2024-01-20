package math

func nextBeautifulNumber(n int) int {
	i := n + 1
	check := func(n int) bool {
		m := map[int]int{}
		for n != 0 {
			mod := n % 10
			if mod > 6 {
				return false
			}
			n /= 10
			m[mod]++
		}
		for k, v := range m {
			if k != v {
				return false
			}
		}
		return true
	}
	for !check(i) {
		i++
	}
	return i
}
