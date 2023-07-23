package array

func largestAltitude(gain []int) int {
	max := 0
	last := 0
	for _, v := range gain {
		last += v
		if last > max {
			max = last
		}
	}
	return max
}
