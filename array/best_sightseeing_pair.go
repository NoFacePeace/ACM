package array

func maxScoreSightseeingPair(values []int) int {
	l := len(values)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return values[0]
	}
	start := 0
	max := values[0]
	for i := 1; i < l; i++ {
		score := values[start] + values[i] + start - i
		if score > max {
			max = score
		}
		if values[i]+i >= values[start]+start {
			start = i
		}
	}
	return max
}
