package math

func checkPowersOfThree(n int) bool {
	p := 1
	arr := []int{}
	for p <= n {
		arr = append(arr, p)
		p *= 3
	}
	var dfs func(i, cnt int) bool
	dfs = func(i, cnt int) bool {
		if cnt > n {
			return false
		}
		if cnt == n {
			return true
		}
		if i >= len(arr) {
			return false
		}
		if dfs(i+1, cnt) {
			return true
		}
		return dfs(i+1, cnt+arr[i])
	}
	return dfs(0, 0)
}
