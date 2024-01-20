package firstsearch

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)
	ad := make([][]int, n)
	degree := make([]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		degree[u]++
		degree[v]++
		ad[u] = append(ad[u], v)
		ad[v] = append(ad[v], u)
	}
	q := []int{}
	dist := make([]int, n)
	visited := make([]int, n)
	for k, v := range degree {
		dist[k] = coins[k]
		if v > 1 {
			continue
		}
		visited[k] = 1
		q = append(q, k)
	}
	new := true
	for new {
		new = false
		tmp := q
		q = nil
		for _, v := range tmp {
			if coins[v] == 1 {
				q = append(q, v)
				continue
			}
			for _, u := range ad[v] {
				if visited[u] == 1 {
					continue
				}
				degree[u]--
				if degree[u] < 2 {
					visited[u] = 1
					q = append(q, u)
					new = true
				}
			}
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, u := range ad[v] {
			if visited[u] == 1 {
				continue
			}
			degree[u]--
			if degree[u] < 2 {
				visited[u] = 1
				q = append(q, u)
			}
			coins[u] += coins[v]
			if coins[v] > 0 {
				dist[u] = max(dist[u], dist[v]+1)
			}
		}
	}
	cnt := 0
	for _, v := range dist {
		if v > 2 {
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return (cnt - 1) * 2
}
