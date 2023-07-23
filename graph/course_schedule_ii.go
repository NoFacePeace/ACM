package graph

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	stack := []int{}
	for _, v := range prerequisites {
		edges[v[0]] = append(edges[v[0]], v[1])
	}
	var dfs func(p int)
	valid := true
	dfs = func(p int) {
		visited[p] = 1
		for _, v := range edges[p] {
			if visited[v] == 1 {
				valid = false
				return
			}
			if visited[v] == 2 {
				continue
			}
			dfs(v)
		}
		if !valid {
			return
		}
		visited[p] = 2
		stack = append(stack, p)
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 2 {
			continue
		}
		dfs(i)
	}
	if valid {
		return stack
	}
	return []int{}
}
