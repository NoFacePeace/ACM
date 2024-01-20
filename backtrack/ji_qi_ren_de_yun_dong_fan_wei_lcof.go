package backtrack

func movingCount(m int, n int, k int) int {
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}
	ans := 0
	var f func(x, y int)
	f = func(x, y int) {
		if x < 0 || x >= m {
			return
		}
		if y < 0 || y >= n {
			return
		}
		if visit[x][y] {
			return
		}
		sum := x%10 + x/10 + y%10 + y/10
		if sum > k {
			return
		}
		visit[x][y] = true
		ans++
		f(x+1, y)
		f(x-1, y)
		f(x, y+1)
		f(x, y-1)
	}
	f(0, 0)
	return ans
}
