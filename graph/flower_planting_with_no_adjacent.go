package graph

func gardenNoAdj(n int, paths [][]int) []int {
	arr := make([][]int, n+1)
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			u, v := paths[i][0], paths[i][1]
			arr[u] = append(arr[u], v)
			arr[v] = append(arr[v], u)
		}
	}
	m := map[int]int{}
	for i := 1; i <= n; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		q := []int{i}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				if _, ok := m[v]; ok {
					continue
				}
				enum := map[int]bool{
					1: true,
					2: true,
					3: true,
					4: true,
				}
				for _, u := range arr[v] {
					if _, ok := m[u]; ok {
						delete(enum, m[u])
					}
					q = append(q, u)
				}
				for k := range enum {
					m[v] = k
					break
				}
			}
		}
	}
	ans := make([]int, n)
	for k, v := range m {
		ans[k-1] = v
	}
	return ans
}
