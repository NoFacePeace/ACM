package array

import "sort"

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	if len(profit) == 0 {
		return 0
	}
	if len(profit) == 1 {
		return profit[0]
	}
	jobs := make([][3]int, len(profit))
	for i := 0; i < len(profit); i++ {
		jobs[i][0] = startTime[i]
		jobs[i][1] = endTime[i]
		jobs[i][2] = profit[i]
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([]int, len(profit)+1)
	for i := 1; i <= len(profit); i++ {
		k := sort.Search(i, func(j int) bool {
			return jobs[j][1] > jobs[i-1][0]
		})
		dp[i] = max(dp[i-1], dp[k]+jobs[i-1][2])
	}
	return dp[len(profit)]
}
