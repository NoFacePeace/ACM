package multiply

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	mod := int(1e9) + 7
	m := map[int]int{}
	for _, v := range arr {
		m[v] = 1
	}
	n := len(arr)
	if n == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += m[arr[i]]
		ans %= mod
		for j := 0; j < i; j++ {
			if m[arr[i]*arr[j]] > 0 {
				m[arr[i]*arr[j]] += m[arr[i]] * m[arr[j]] * 2
			}
		}
		if m[arr[i]*arr[i]] > 0 {
			m[arr[i]*arr[i]] += m[arr[i]] * m[arr[i]]
		}
	}
	return ans
}
