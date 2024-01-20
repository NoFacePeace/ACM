package dfa

import "math"

func strToInt(str string) int {
	isInt := func(str string) (bool, int, int, int) {
		n := len(str)
		if n == 0 {
			return false, 0, 0, 0
		}
		i := 0
		for i < n {
			if str[i] != ' ' {
				break
			}
			i++
		}
		if i == n {
			return false, 0, 0, 0
		}
		sign := 1
		if str[i] == '-' {
			sign = -1
			i++
		} else if str[i] == '+' {
			i++
		}
		if i == n {
			return false, 0, 0, 0
		}
		if str[i] < '0' || str[i] > '9' {
			return false, 0, 0, 0
		}
		left, right := i, i
		i++
		for i < n {
			if str[i] < '0' || str[i] > '9' {
				break
			}
			right = i
			i++
		}
		return true, sign, left, right
	}
	ok, sign, left, right := isInt(str)
	if !ok {
		return 0
	}
	num := 0
	for i := left; i <= right; i++ {
		num *= 10
		num += int(str[i]-'0') * sign
		if num > math.MaxInt32 {
			return math.MaxInt32
		}
		if num < math.MinInt32 {
			return math.MinInt32
		}
	}
	return num
}
