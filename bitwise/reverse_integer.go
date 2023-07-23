package bitwise

func reverse(x int) int {
	min := ^(1 << 31)
	max := 1 << 31
	if x == min {
		return 0
	}
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}
	ans := 0
	multiple := 1
	for x > 0 {
		mod := x % 10
		if (max-mod)/multiple < ans {
			return 0
		}
		ans *= multiple
		ans += mod
		x /= 10
		multiple = 10
	}
	if isNegative {
		ans = -ans
	}
	return ans
}
