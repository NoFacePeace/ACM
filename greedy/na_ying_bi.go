package greedy

func minCount(coins []int) int {
	cnt := 0
	for _, v := range coins {
		cnt += v / 2
		if v%2 != 0 {
			cnt++
		}
	}
	return cnt
}
