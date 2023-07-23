package binarysearch

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	n := len(nums1)
	if n == 0 {
		return nil
	}
	m := len(nums2)
	if m == 0 {
		return nil
	}
	left := nums1[0] + nums2[0]
	right := nums1[n-1] + nums2[m-1]
	for left < right {
		mid := left + (right-left)/2
		i := n - 1
		j := 0
		cnt := 0
		for i >= 0 && j < m {
			if nums1[i]+nums2[j] > mid {
				i--
			} else {
				cnt += i + 1
				j++
			}
		}
		if cnt >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	ans := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i]+nums2[j] >= left {
				break
			}
			ans = append(ans, []int{nums1[i], nums2[j]})
			if len(ans) == k {
				return ans
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i]+nums2[j] > left {
				break
			}
			if nums1[i]+nums2[j] < left {
				continue
			}
			ans = append(ans, []int{nums1[i], nums2[j]})
			if len(ans) == k {
				return ans
			}
		}
	}
	return ans
}
