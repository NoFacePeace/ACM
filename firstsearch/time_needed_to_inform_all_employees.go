package firstsearch

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	ad := make([][]int, n)
	for k, v := range manager {
		if v == -1 {
			continue
		}
		ad[v] = append(ad[v], k)
	}
	ans := 0
	var dfs func(id, cost int)
	dfs = func(id, cost int) {
		if len(ad[id]) == 0 {
			if cost > ans {
				ans = cost
			}
			return
		}
		for _, v := range ad[id] {
			dfs(v, cost+informTime[id])
		}
	}
	dfs(headID, 0)
	return ans
}
