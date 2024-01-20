package math

func sumOfMultiples(n int) int {
	tree := 3
	five := 5
	seven := 7
	ans := 0
	min := 0
	for min <= n {
		ans += min
		min = func(a, b, c int) int {
			if a <= b && a <= c {
				return a
			}
			if b <= a && b <= c {
				return b
			}
			return c
		}(tree, five, seven)
		if min == tree {
			tree += 3
		}
		if min == five {
			five += 5
		}
		if min == seven {
			seven += 7
		}
	}
	return ans
}
