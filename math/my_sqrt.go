package math

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	left := 0
	right := x
	for left < right {
		mid := (left + right) / 2
		if mid == left || mid == right {
			return mid
		}
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			left = mid
			continue
		}
		right = mid
	}
	return left
}
