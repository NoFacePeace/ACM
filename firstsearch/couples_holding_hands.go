package firstsearch

func minSwapsCouples(row []int) int {
	n := len(row)
	ad := make([][]int, n/2)
	for i := 0; i < n; i += 2 {
		u, v := row[i]/2, row[i+1]/2
		ad[u] = append(ad[u], v)
		ad[v] = append(ad[v], u)
	}
	cnt := 0
	visited := make([]bool, n/2)
	var dfs func(idx int)
	dfs = func(idx int) {
		if visited[idx] {
			return
		}
		visited[idx] = true
		for _, v := range ad[idx] {
			dfs(v)
		}
	}
	for i := 0; i < n/2; i++ {
		if visited[i] {
			continue
		}
		cnt++
		dfs(i)
	}
	return n/2 - cnt
}
