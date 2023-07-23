package firstsearch

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	ad := map[int][]int{}
	for _, e := range edges {
		u, v := e[0], e[1]
		ad[u] = append(ad[u], v)
		ad[v] = append(ad[v], u)
	}
	comb := func(n int) int {
		num := 1
		for i := n; i > n-2; i-- {
			num *= i
		}
		return num / 2
	}
	ans := make([]int, n-1)
	var f func(u, fa, cnt int)
	f = func(u, fa, cnt int) {
		l := len(ad[u])
		if fa != 0 {
			l -= 1
		}
		if l == 0 {
			return
		}
		for i := 0; i <= cnt; i++ {
			ans[i] += l
		}
		inc := comb(l)
		if inc != 0 {
			ans[1] += inc
			for i := 1; i <= cnt; i++ {
				ans[cnt] += inc
			}
		}
		for _, v := range ad[u] {
			if v == fa {
				continue
			}
			f(v, u, cnt+1)
		}
	}
	f(1, 0, 0)
	return ans
}
