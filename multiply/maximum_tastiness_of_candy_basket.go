package multiply

import "sort"

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	left := 0
	n := len(price)
	right := price[n-1] - price[0]
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		l := 0
		cnt := 0
		for i := 1; i < n; i++ {
			dist := price[i] - price[l]
			if dist < mid {
				continue
			}
			l = i
			cnt++
		}
		if cnt < k-1 {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}
