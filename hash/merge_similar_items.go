package hash

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, v := range items1 {
		m[v[0]] += v[1]
	}
	for _, v := range items2 {
		m[v[0]] += v[1]
	}
	ans := [][]int{}
	for k, v := range m {
		ans = append(ans, []int{k, v})
	}
	sort.Slice(ans, func(a, b int) bool {
		if ans[a][0] < ans[b][0] {
			return true
		}
		return false
	})
	return ans
}
