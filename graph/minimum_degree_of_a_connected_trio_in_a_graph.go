package graph

func minTrioDegree(n int, edges [][]int) int {
	ad := make([][]bool, n)
	for i := 0; i < n; i++ {
		ad[i] = make([]bool, n)
	}
	m := map[int]int{}
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		ad[u][v] = true
		ad[v][u] = true
		m[u]++
		m[v]++
	}
	max := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !ad[i][j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if !ad[j][k] {
					continue
				}
				if !ad[i][k] {
					continue
				}
				degree := m[i] + m[j] + m[k] - 6
				if max == -1 {
					max = degree
					continue
				}
				if max > degree {
					max = degree
				}
			}
		}
	}
	return max
}
