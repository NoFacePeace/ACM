package firstsearch

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := make([][]int, numCourses)
	for _, v := range prerequisites {
		u, v := v[0], v[1]
		g[u] = append(g[u], v)
	}
	ans := []int{}
	visit := map[int]bool{}
	finish := map[int]bool{}
	var f func(num int) bool
	f = func(num int) bool {
		if visit[num] {
			return false
		}
		visit[num] = true
		for i := 0; i < len(g[num]); i++ {
			if finish[g[num][i]] {
				continue
			}
			if !f(g[num][i]) {
				return false
			}
		}
		finish[num] = true
		ans = append(ans, num)
		return true
	}
	num := 0
	for len(visit) < numCourses {
		if finish[num] {
			num++
			continue
		}
		if !f(num) {
			return []int{}
		}
		num++
	}
	return ans
}
