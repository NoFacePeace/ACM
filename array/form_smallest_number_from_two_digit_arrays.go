package array

import "math"

func minNumber(nums1 []int, nums2 []int) int {
	m := map[int]bool{}
	mn1 := math.MaxInt
	for _, v := range nums1 {
		if v < mn1 {
			mn1 = v
		}
		m[v] = true
	}
	mn2 := math.MaxInt
	mn3 := math.MaxInt
	for _, v := range nums2 {
		if v < mn2 {
			mn2 = v
		}
		if m[v] {
			if v < mn3 {
				mn3 = v
			}
		}
	}

	if mn1*10+mn2 > mn1+mn2*10 {
		mn1 = mn1 + mn2*10
	} else {
		mn1 = mn1*10 + mn2
	}
	if mn1 < mn3 {
		return mn1
	}
	return mn3
}
