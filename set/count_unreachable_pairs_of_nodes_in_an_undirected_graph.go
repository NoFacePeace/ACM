package set

import . "github.com/NoFacePeace/ACM/packages"

func countPairs(n int, edges [][]int) int64 {
	set := NewUnionFind(n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		set.Union(u, v)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += n - set.GetSize(set.Find(i))
	}
	return int64(ans / 2)
}
