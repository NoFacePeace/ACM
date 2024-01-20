package graph

import (
	"math"
)

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	ad := make([][]int, n)
	for i := 0; i < n; i++ {
		ad[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ad[i][j] = math.MaxInt / 2
		}
	}
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		ad[u][v] = w
		ad[v][u] = w
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 0; i < n; i++ {
		ad[i][i] = 0
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				ad[j][k] = min(ad[j][k], ad[j][i]+ad[i][k])
			}
		}
	}
	ans := math.MaxInt
	idx := -1
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if ad[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= ans {
			ans = cnt
			idx = i
		}
	}
	return idx
}
