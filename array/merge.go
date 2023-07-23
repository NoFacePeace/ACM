package array

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := []int{}
	tmp = append(tmp, nums1...)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if tmp[i] > nums2[j] {
			nums1[k] = nums2[j]
			k += 1
			j += 1
			continue
		}
		nums1[k] = tmp[i]
		k += 1
		i += 1
	}
	for i < m {
		nums1[k] = tmp[i]
		k += 1
		i += 1
	}
	for j < n {
		nums1[k] = nums2[j]
		k += 1
		j += 1
	}
}
