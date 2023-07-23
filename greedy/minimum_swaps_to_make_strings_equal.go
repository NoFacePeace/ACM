package greedy

func minimumSwap(s1 string, s2 string) int {
	l := len(s1)
	x := 0
	y := 0
	for i := 0; i < l; i++ {
		if s1[i] == s2[i] {
			continue
		}
		if s1[i] == 'x' {
			x++
		}
		if s1[i] == 'y' {
			y++
		}
	}
	ans := 0
	ans += x / 2
	ans += y / 2
	x = x % 2
	y = y % 2
	if x != y {
		return -1
	}
	ans += x + y
	return ans
}
