package slidingwindow

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	if n < k {
		return 0
	}
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}
	max := sum
	j := n - 1
	for i := k - 1; i >= 0; i-- {
		sum -= cardPoints[i]
		sum += cardPoints[j]
		j--
		if sum > max {
			max = sum
		}
	}
	return max
}
