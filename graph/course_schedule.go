package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	stack := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		p := prerequisites[i]
		edges[p[0]] = append(edges[p[0]], p[1])
	}
	valid := true
	var dfs func(n int)
	dfs = func(n int) {
		visited[n] = 1
		for _, v := range edges[n] {
			if visited[v] == 1 {
				valid = false
				return
			}
			if visited[v] == 2 {
				continue
			}
			dfs(v)
			if !valid {
				return
			}
		}
		visited[n] = 2
		stack = append(stack, n)
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] != 0 {
			continue
		}
		dfs(i)
	}
	return valid
}
