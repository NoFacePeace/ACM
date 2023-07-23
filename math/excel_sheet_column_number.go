package math

func titleToNumber(columnTitle string) int {
	if len(columnTitle) == 0 {
		return 0
	}
	ans := 0
	tmp := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		diff := int(columnTitle[i] - 'A' + 1)
		ans += diff * tmp
		tmp *= 26
	}
	return ans
}
