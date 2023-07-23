package array

func maximumValue(strs []string) int {
	ans := 0
	isNum := func(s string) bool {
		for i := 0; i < len(s); i++ {
			if s[i] < '0' {
				return false
			}
			if s[i] > '9' {
				return false
			}
		}
		return true
	}
	convert := func(s string) int {
		num := 0
		for i := 0; i < len(s); i++ {
			num *= 10
			num += int(s[i] - '0')
		}
		return num
	}
	for _, s := range strs {
		var val int
		if !isNum(s) {
			val = len(s)
		} else {
			val = convert(s)
		}
		if val > ans {
			ans = val
		}
	}
	return ans
}
