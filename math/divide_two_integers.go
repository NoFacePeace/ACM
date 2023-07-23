package math

import (
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	isNegative := false
	if dividend < 0 {
		isNegative = !isNegative
		dividend = -dividend
	}
	if divisor < 0 {
		isNegative = !isNegative
		divisor = -divisor
	}
	count := quickAdd(divisor, dividend)
	if isNegative {
		return -count
	}
	return count
}

func quickAdd(x, y int) int {
	x1 := x
	if x > y {
		return 0
	}
	if x == y {
		return 1
	}
	count := 1
	for x+x <= y {
		count += count
		x += x
	}
	if x == y {
		return count
	}
	return count + quickAdd(x1, y-x)
}
