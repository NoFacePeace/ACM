package dp

import "sort"

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(a, b int) bool {
		return rides[a][1] < rides[b][1]
	})
	dp := make([]int, n+1)
	j := 0
	i := 1
	for j < len(rides) {
		if i != rides[j][1] {
			dp[i] = max(dp[i], dp[i-1])
			i++
			continue
		}
		start, end, tip := rides[j][0], rides[j][1], rides[j][2]
		j++
		dp[i] = max(dp[i], dp[i-1])
		dp[i] = max(dp[i], end-start+tip+dp[start])
	}
	return int64(dp[i])
}
