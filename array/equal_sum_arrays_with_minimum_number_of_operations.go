package array

import "sort"

func minOperations1(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	min1 := l1
	max1 := l1 * 6
	min2 := l2
	max2 := l2 * 6
	if min1 > max2 {
		return -1
	}
	if min2 > max1 {
		return -1
	}
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	c1 := 0
	for _, v := range nums1 {
		c1 += v
	}
	c2 := 0
	for _, v := range nums2 {
		c2 += v
	}
	if c1 < c2 {
		c1, c2 = c2, c1
		nums1, nums2 = nums2, nums1
	}
	i1 := len(nums1) - 1
	i2 := 0
	ans := 0
	for c1 > c2 {
		if i1 >= len(nums1) {
			c2 += 6 - nums2[i2]
			i2++
			ans++
			continue
		}
		if i2 >= len(nums2) {
			c1 -= nums1[i1] - 1
			ans++
			i1--
			continue
		}
		if nums1[i1]-1 > 6-nums2[i2] {
			c1 -= nums1[i1] - 1
			ans++
			i1--
			continue
		}
		c2 += 6 - nums2[i2]
		i2++
		ans++
		continue
	}
	return ans
}
