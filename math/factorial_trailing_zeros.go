package math

func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5
		n = n / 5
	}
	return ans
}
