package firstsearch

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans := [][]int{}
	var f func(arr [][]int, u int)
	f = func(arr [][]int, u int) {
		if u == n-1 {
			tmp := make([][]int, len(arr))
			for i := 0; i < len(arr); i++ {
				tmp[i] = append([]int{}, arr[i]...)
			}
			ans = append(ans, tmp...)
			return
		}
		for _, v := range graph[u] {
			tmp := append([][]int{}, arr...)
			for i := 0; i < len(tmp); i++ {
				tmp[i] = append(tmp[i], v)
			}
			f(tmp, v)
		}
	}
	f([][]int{{0}}, 0)
	return ans
}
