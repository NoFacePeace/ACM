package pointer

import "sort"

func countPairs1(n int, edges [][]int, queries []int) []int {
	cnt := make([]int, n)
	m := map[int]int{}
	for _, e := range edges {
		v1, v2 := e[0]-1, e[1]-1
		if v1 > v2 {
			v1, v2 = v2, v1
		}
		cnt[v1]++
		cnt[v2]++
		m[v1*n+v2]++
	}
	st := []int{}
	st = append(st, cnt...)
	sort.Ints(st)
	ans := []int{}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _, q := range queries {
		total := 0
		for i, j := 0, n-1; i < n; i++ {
			for j > i && st[i]+st[j] > q {
				j--
			}
			total += n - 1 - max(i, j)
		}
		for k, v := range m {
			v1, v2 := k/n, k%n
			if cnt[v1]+cnt[v2] > q && cnt[v1]+cnt[v2]-v <= q {
				total--
			}
		}
		ans = append(ans, total)
	}
	return ans
}
