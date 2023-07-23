package str

func secondHighest(s string) int {
	max := -1
	ans := -1
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			continue
		}
		num := int(s[i] - '0')
		if max == -1 {
			max = num
			continue
		}
		if num == max {
			continue
		}
		if num > max {
			max, ans = num, max
			continue
		}
		if num > ans {
			ans = num
		}
	}
	return ans
}
