package firstsearch

func minimumFuelCost(roads [][]int, seats int) int64 {
	ad := map[int][]int{}
	for _, road := range roads {
		u, v := road[0], road[1]
		ad[u] = append(ad[u], v)
		ad[v] = append(ad[v], u)
	}
	cost := 0
	visited := map[int]bool{}
	var dfs func(u int) int
	dfs = func(u int) int {
		visited[u] = true
		cnt := 1
		for _, v := range ad[u] {
			if visited[v] {
				continue
			}
			cnt += dfs(v)
		}
		if u == 0 {
			return cnt
		}
		cost += (cnt + seats - 1) / seats
		return cnt
	}
	dfs(0)
	return int64(cost)
}
