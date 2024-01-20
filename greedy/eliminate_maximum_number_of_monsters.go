package greedy

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		t := (dist[i] - 1) / speed[i]
		if t >= n {
			continue
		}
		cnt[t]++
	}
	kill := 0
	i := 0
	for i < n {
		kill++
		kill -= cnt[i]
		if kill < 0 {
			break
		}
		i++
	}
	if i == n {
		return i
	}
	return i + 1
}
