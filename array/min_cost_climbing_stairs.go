package array

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	if l == 0 {
		return 0
	}
	if l == 1 {
		return cost[0]
	}
	if l == 2 {
		return min(cost[0], cost[1])
	}
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[l-1], cost[l-2])
}
