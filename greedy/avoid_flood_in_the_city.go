package greedy

func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	free := []int{}
	m := map[int]int{}
	for k, v := range rains {
		if v == 0 {
			ans[k] = 1
			free = append(free, k)
			continue
		}
		if _, ok := m[v]; !ok {
			m[v] = k
			ans[k] = -1
			continue
		}
		status := false
		i := 0
		for i < len(free) {
			if free[i] < m[v] {
				i++
				continue
			}
			ans[k] = -1
			status = true
			ans[free[i]] = v
			m[v] = k
			tmp := append([]int{}, free[:i]...)
			tmp = append(tmp, free[i+1:]...)
			free = tmp
			break
		}
		if !status {
			return []int{}
		}
	}
	return ans
}
