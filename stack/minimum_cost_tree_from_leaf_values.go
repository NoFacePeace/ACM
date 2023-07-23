package stack

func mctFromLeafValues(arr []int) int {
	st := []int{}
	ans := 0
	for _, v := range arr {
		n := len(st)
		if n == 0 {
			st = append(st, v)
			continue
		}
		if v <= st[n-1] {
			st = append(st, v)
			continue
		}
		for n > 0 && st[n-1] < v {
			if n == 1 || st[n-2] > v {
				ans += st[n-1] * v
			} else {
				ans += st[n-1] * st[n-2]
			}
			st = st[:n-1]
			n--
		}
		st = append(st, v)
	}
	for len(st) > 1 {
		n := len(st)
		ans += st[n-1] * st[n-2]
		st = st[:n-1]
	}
	return ans
}
