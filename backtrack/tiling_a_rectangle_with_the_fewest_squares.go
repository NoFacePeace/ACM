package backtrack

func tilingRectangle(n int, m int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}
	allow := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if arr[x+i][y+j] == 1 {
					return false
				}
			}
		}
		return true
	}
	fill := func(x, y, k, val int) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				arr[i+x][j+y] = val
			}
		}
	}
	ans := n * m
	var dfs func(x, y, cnt int)
	dfs = func(x, y, cnt int) {
		if cnt >= ans {
			return
		}
		if x >= n {
			ans = cnt
			return
		}
		if y >= m {
			dfs(x+1, 0, cnt)
			return
		}
		if arr[x][y] == 1 {
			dfs(x, y+1, cnt)
			return
		}
		k := m - y
		if k > n-x {
			k = n - x
		}
		for k > 0 {
			if allow(x, y, k) {
				break
			}
			k--
		}
		for i := k; i >= 0; i-- {
			fill(x, y, i, 1)
			dfs(x, y+i, cnt+1)
			fill(x, y, i, 0)
		}
	}
	dfs(0, 0, 0)
	return ans
}
