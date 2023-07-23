package graph

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	for i := 0; i < len(edges); i++ {
		v, e, w := edges[i][0], edges[i][1], edges[i][2]
		arr[v][e] = w
		arr[e][v] = w
	}
	var dfs func(e, t int, q [][]int)
	m := map[int]bool{}
	ok := false
	dfs = func(e, t int, q [][]int) {
		if e == destination {
			if t < 0 {
				return
			}
			if t > 0 && len(q) == 0 {
				edges = [][]int{}
				return
			}
			ok = true
			for i := 0; i < len(q); i++ {
				v, e := q[i][0], q[i][1]
				if i == len(q)-1 {
					arr[v][e] = t
					arr[e][v] = t
					continue
				}
				arr[v][e] = 1
				arr[e][v] = 1
				t--
			}
			return
		}
		for i := 0; i < len(arr[e]); i++ {
			if arr[e][i] == 0 {
				continue
			}
			if m[i] {
				continue
			}
			m[i] = true
			tmp := t
			if arr[e][i] == -1 {
				q = append(q, []int{e, i})
			} else {
				tmp = t - arr[e][i]
			}
			dfs(i, tmp, q)
			m[i] = false
			if arr[e][i] == -1 {
				q = q[:len(q)-1]
			}
		}
	}
	m[source] = true
	dfs(source, target, [][]int{})
	if !ok {
		return [][]int{}
	}
	for i := 0; i < len(edges); i++ {
		v, e, w := edges[i][0], edges[i][1], edges[i][2]
		if w != -1 {
			continue
		}
		if arr[v][e] == -1 {
			edges[i][2] = 1
			continue
		}
		edges[i][2] = arr[v][e]
	}
	return edges
}
