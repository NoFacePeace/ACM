package math

func alternateDigitSum(n int) int {
	signal := 1
	ans := 0
	for n != 0 {
		val := n % 10 * signal
		signal = -signal
		n /= 10
		ans += val
	}
	if signal == 1 {
		ans = -ans
	}
	return ans
}
