package dp

func numTrees(n int) int {
	var dfs func(n int) int
	m := map[int]int{}
	m[0] = 0
	m[1] = 1
	dfs = func(n int) int {
		if n == 0 {
			return 1
		}
		if n == 1 {
			return 1
		}
		if m[n] != 0 {
			return m[n]
		}
		cnt := 0
		for i := 1; i <= n; i++ {
			cnt += dfs(i-1) * dfs(n-i)
		}
		m[n] = cnt
		return cnt
	}
	return dfs(n)
}
