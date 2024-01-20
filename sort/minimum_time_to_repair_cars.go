package sort

import (
	"math"
	"sort"
)

func repairCars(ranks []int, cars int) int64 {
	n := len(ranks)
	if n == 0 {
		return 0
	}
	sort.Ints(ranks)
	rank := ranks[0]
	time := rank * cars * cars
	l := 1
	r := time
	for l < r {
		mid := (l + r) / 2
		cnt := 0
		for j := 0; j < n; j++ {
			cnt += int(math.Sqrt(float64(mid / ranks[j])))
			if cnt > cars {
				break
			}
		}
		if cnt >= cars {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return int64(r)
}
