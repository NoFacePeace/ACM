package hash

func circularGameLosers(n int, k int) []int {
	m := map[int]int{}
	i := 0
	cnt := 1
	for {
		m[i]++
		if m[i] == 2 {
			break
		}
		i = (i + cnt*k) % n
		cnt++
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		if _, ok := m[i]; !ok {
			ans = append(ans, i+1)
		}
	}
	return ans
}
