package firstsearch

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	stop := map[int][]int{}
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			stop[routes[i][j]] = append(stop[routes[i][j]], i)
		}
	}
	q := []int{target}
	visit := map[int]bool{}
	walk := map[int]bool{}
	m := map[int]int{}
	m[target] = 0
	cnt := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if visit[v] {
				continue
			}
			visit[v] = true
			for _, r := range stop[v] {
				if walk[r] {
					continue
				}
				walk[r] = true
				for _, s := range routes[r] {
					if visit[s] {
						continue
					}
					if m[s] == 0 {
						m[s] = cnt + 1
					}
					if !visit[s] {
						q = append(q, s)
					}
				}
			}
		}
		cnt++
	}
	if m[source] == 0 {
		return -1
	}
	return m[source]
}
