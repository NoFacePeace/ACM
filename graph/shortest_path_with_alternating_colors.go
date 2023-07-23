package graph

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type pair struct {
		edge  int
		color int
	}
	g := make([][]pair, n)
	for _, e := range redEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 0})
	}
	for _, e := range blueEdges {
		g[e[0]] = append(g[e[0]], pair{e[1], 1})
	}
	q := []pair{
		pair{0, 0},
		pair{0, 1},
	}
	visit := make([][2]bool, n)
	level := 0
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		ans[i] = -1
	}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			if ans[p.edge] < 0 {
				ans[p.edge] = level
			}
			for _, to := range g[p.edge] {
				if !visit[to.edge][to.color] && to.color != p.color {
					q = append(q, to)
					visit[to.edge][to.color] = true
				}
			}
		}
		level++
	}
	return ans
}
