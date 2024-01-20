package firstsearch

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	n := len(parents)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}
	node := -1
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			node = i
		}
	}
	min := 1
	ad := make([][]int, n)
	for i := 0; i < n; i++ {
		p := parents[i]
		if p == -1 {
			continue
		}
		ad[p] = append(ad[p], i)
	}
	dna := map[int]bool{}
	visited := map[int]bool{}
	var dfs func(u int)
	dfs = func(u int) {
		if visited[u] {
			return
		}
		visited[u] = true
		dna[nums[u]] = true
		for _, v := range ad[u] {
			dfs(v)
		}
	}
	for node != -1 {
		dfs(node)
		for dna[min] {
			min++
		}
		ans[node] = min
		node = parents[node]
	}
	return ans
}
