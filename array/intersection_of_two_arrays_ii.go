package array

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	ans := []int{}
	for k := range nums1 {
		m[nums1[k]]++
	}
	for k := range nums2 {
		m[nums2[k]]--
		if m[nums2[k]] >= 0 {
			ans = append(ans, nums2[k])
		}
	}
	return ans
}
