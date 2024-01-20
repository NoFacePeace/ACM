package firstsearch

func canFinish(numCourses int, prerequisites [][]int) bool {
	ad := make([][]int, numCourses)
	for _, p := range prerequisites {
		u, v := p[0], p[1]
		ad[v] = append(ad[v], u)
	}
	visited := make([]int, numCourses)
	valid := true
	var dfs func(n int)
	dfs = func(u int) {
		visited[u] = 1
		for _, v := range ad[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
				continue
			}
			if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}
