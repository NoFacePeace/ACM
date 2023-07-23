package array

import "math"

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	if len(customers) == 0 {
		return 0
	}
	i := 0
	wait := customers[i]
	total := 0
	max := math.MinInt
	index := -1
	for wait > 0 || i < len(customers) {
		cnt := wait
		if cnt > 4 {
			wait = cnt - 4
			cnt = 4
		} else {
			wait = 0
		}
		total += cnt
		cost := total*boardingCost - (i+1)*runningCost
		if cost > max {
			max = cost
			index = i + 1
		}
		i++
		if i < len(customers) {
			wait += customers[i]
		}
	}
	if max <= 0 {
		return -1
	}
	return index
}
