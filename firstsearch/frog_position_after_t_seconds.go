package firstsearch

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	ad := make([][]int, n+1)
	for i := 0; i < len(edges); i++ {
		v, u := edges[i][0], edges[i][1]
		ad[v] = append(ad[v], u)
		ad[u] = append(ad[u], v)
	}
	q := []int{1}
	s := 0
	m := map[int]float64{}
	path := map[int]int{}
	m[1] = 1
	status := false
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			l := float64(len(ad[v]))
			if v != 1 {
				l = l - 1
			}
			for _, vv := range ad[v] {
				if vv == path[v] {
					continue
				}
				m[vv] = 1 / l * m[v]
				path[vv] = v
				q = append(q, vv)
				if vv == target {
					status = true
					break
				}
			}
			if status {
				break
			}
		}
		s++
		if status {
			break
		}
		if s == t {
			break
		}
	}
	if s < t && len(ad[target]) > 1 {
		return 0
	}
	if target == 1 && len(ad[target]) > 0 {
		return 0
	}
	return m[target]
}
