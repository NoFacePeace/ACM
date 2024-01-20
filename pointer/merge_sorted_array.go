package pointer

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n
	l--
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[l] = nums1[m]
			l--
			m--
		} else {
			nums1[l] = nums2[n]
			n--
			l--
		}
	}
	for n >= 0 {
		nums1[l] = nums2[n]
		n--
		l--
	}
}
