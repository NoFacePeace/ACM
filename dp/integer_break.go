package dp

func integerBreak(n int) int {
	m := map[int]int{}
	m[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if m[i-j]*j > m[i] {
				m[i] = m[i-j] * j
			}
			if j*(i-j) > m[i] {
				m[i] = j * (i - j)
			}
		}
	}
	return m[n]
}
