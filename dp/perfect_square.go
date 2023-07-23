package dp

func numSquares(n int) int {
	m := map[int]int{}
	var f func(i int) int
	f = func(n int) int {
		if m[n] != 0 {
			return m[n]
		}
		if n == 1 {
			m[n] = 1
			return 1
		}
		i := 0
		for ; i < n/2; i++ {
			if (i+1)*(i+1) > n {
				break
			}
		}
		if i*i == n {
			m[n] = 1
			return 1
		}
		min := n
		for i := 1; i*i < n; i++ {
			if f(n-i*i)+1 < min {
				min = f(n-i*i) + 1
			}
		}
		m[n] = min
		return min
	}
	return f(n)
}
