package dp

func sumOfDistancesInTree(n int, edges [][]int) []int {
	ad := make([][]int, n)
	for _, e := range edges {
		v1, v2 := e[0], e[1]
		ad[v1] = append(ad[v1], v2)
		ad[v2] = append(ad[v2], v1)
	}
	ans := make([]int, n)
	cnt := make([]int, n)
	m := map[int]bool{}
	var dfs func(num int) (int, int)
	dfs = func(num int) (int, int) {
		m[num] = true
		for _, v := range ad[num] {
			if v == num {
				continue
			}
			if m[v] {
				continue
			}
			val, c := dfs(v)
			ans[num] += val
			cnt[num] += c
		}
		ans[num] += cnt[num]
		cnt[num] += 1
		return ans[num], cnt[num]
	}
	dfs(0)
	q := []int{0}
	m = map[int]bool{}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			m[v] = true
			for _, e := range ad[v] {
				if m[e] {
					continue
				}
				ans[e] = ans[v] - cnt[e] + 1 + (n - cnt[e] - 1)
				q = append(q, e)
			}
		}
	}
	return ans
}
