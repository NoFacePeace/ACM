package dp

func minimumTime(n int, relations [][]int, time []int) int {
	ad := make([][]int, n+1)
	for _, v := range relations {
		ad[v[1]] = append(ad[v[1]], v[0])
	}
	res := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m := map[int]int{}
	var dp func(idx int) int
	dp = func(idx int) int {
		if m[idx] != 0 {
			return m[idx]
		}
		month := 0
		for _, v := range ad[idx] {
			month = max(month, dp(v))
		}
		month += time[idx-1]
		m[idx] = month
		return m[idx]
	}
	for i := 1; i <= n; i++ {
		res = max(res, dp(i))
	}
	return res
}
