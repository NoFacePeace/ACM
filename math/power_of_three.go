package math

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	multiple := 3
	for multiple <= n {
		if n == multiple {
			return true
		}
		multiple *= 3
	}
	return false
}
