package firstsearch

func minReorder(n int, connections [][]int) int {
	ad := make([][]int, n)
	m := map[int]bool{}
	for _, c := range connections {
		u, v := c[0], c[1]
		ad[u] = append(ad[u], v)
		ad[v] = append(ad[v], u)
		m[u*n+v] = true
	}
	cnt := 0
	visited := make([]bool, n)
	var dfs func(u int)
	dfs = func(u int) {
		visited[u] = true
		for _, v := range ad[u] {
			if visited[v] {
				continue
			}
			if m[u*n+v] {
				cnt++
			}
			dfs(v)
		}
	}
	dfs(0)
	return cnt
}
