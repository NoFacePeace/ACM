package pointer

func maxDistToClosest(seats []int) int {
	max := 0
	n := len(seats)
	l := 0
	for l < n && seats[l] == 0 {
		l++
		max++
	}
	r := l + 1
	for r < n {
		if seats[r] == 0 {
			r++
			continue
		}
		mid := (r - l) / 2
		if mid > max {
			max = mid
		}
		l = r
		r = r + 1
	}
	if seats[n-1] == 0 {
		if n-l-1 > max {
			max = n - l - 1
		}
	}
	return max
}
