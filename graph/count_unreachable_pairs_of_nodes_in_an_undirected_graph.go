package graph

func countPairs(n int, edges [][]int) int64 {
	ad := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		ad[a] = append(ad[a], b)
		ad[b] = append(ad[b], a)
	}
	visited := map[int]bool{}
	cnts := 0
	cnt := 0
	ans := 0
	var dfs func(u int)
	dfs = func(u int) {
		for _, v := range ad[u] {
			if visited[v] {
				continue
			}
			visited[v] = true
			cnt++
			dfs(v)
		}
	}
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		cnt = 1
		visited[i] = true
		dfs(i)
		ans += cnts * cnt
		cnts += cnt
	}
	return int64(ans)
}
