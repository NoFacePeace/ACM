package math

import "sort"

func permutation(s string) []string {
	t := []byte(s)
	ans := []string{}
	sort.Slice(t, func(a, b int) bool {
		return t[a] < t[b]
	})
	reverse := func(a []byte) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	next := func(nums []byte) bool {
		n := len(nums)
		i := n - 2
		for i >= 0 && nums[i] >= nums[i+1] {
			i--
		}
		if i < 0 {
			return false
		}
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		reverse(nums[i+1:])
		return true
	}
	for {
		ans = append(ans, string(t))
		if !next(t) {
			break
		}
	}
	return ans
}
