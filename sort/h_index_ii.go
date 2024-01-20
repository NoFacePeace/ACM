package sort

func hIndex(citations []int) int {
	n := len(citations)
	l := 0
	r := n
	for l < r {
		m := (l + r) / 2
		if citations[m] >= n-m {
			r = m
		} else {
			l = m + 1
		}
	}

	return n - l
}
