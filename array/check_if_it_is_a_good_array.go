package array

import "sort"

func isGoodArray(nums []int) bool {
	type pair struct {
		val  int
		base int
	}
	q := []pair{}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	for i := 0; i < len(nums); i++ {
		q = append(q, pair{nums[i], nums[i]})
	}
	for len(q) > 1 {
		if q[0].val+1 == q[1].val {
			return true
		}
		base := q[0].base
		val := q[0].val + base
		q[0] = pair{val, base}
		for i := 0; i < len(q)-1; i++ {
			if q[i].val > q[i+1].val {
				q[i], q[i+1] = q[i+1], q[i]
			}
		}
	}
	return false
}
