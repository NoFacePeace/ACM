package math

func numPairsDivisibleBy60(time []int) int {
	m := map[int]int{}
	for _, v := range time {
		m[v%60]++
	}
	ans := 0
	for k, v := range m {
		if k > 30 {
			continue
		}
		if k == 0 || k == 30 {
			for i := 1; i < v; i++ {
				ans += i
			}
			continue
		}
		if _, ok := m[60-k]; !ok {
			continue
		}
		ans += m[k] * m[60-k]
	}
	return ans
}
