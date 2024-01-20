package sort

import (
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(a, b int) bool {
		s1 := strconv.Itoa(nums[a])
		s2 := strconv.Itoa(nums[b])
		l1 := len(s1)
		l2 := len(s2)
		i, j := 0, 0
		for i < l1 || j < l2 {
			if i >= l1 {
				i = 0
			}
			if j >= l2 {
				j = 0
			}
			if s1[i] < s2[j] {
				return true
			}
			if s1[i] > s2[j] {
				return false
			}
			i++
			j++
		}
		return true
	})
	str := ""
	for _, v := range nums {
		str += strconv.Itoa(v)
	}
	return str
}
