package str

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	ans := ""
	if numerator < 0 != (denominator < 0) {
		ans = ans + "-"
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	ans += strconv.Itoa(numerator / denominator)
	ans += "."
	numerator = numerator % denominator
	m := map[int]int{}
	i := len(ans)
	for numerator != 0 && m[numerator] == 0 {
		m[numerator] = i
		i++
		numerator *= 10
		ans += strconv.Itoa(numerator / denominator)
		numerator = numerator % denominator
	}
	if numerator != 0 {
		return ans[:m[numerator]] + "(" + ans[m[numerator]:] + ")"
	}
	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
