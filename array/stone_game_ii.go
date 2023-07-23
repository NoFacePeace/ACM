package array

import "math"

func stoneGameII(piles []int) int {
	l := len(piles)
	if l == 0 {
		return 0
	}
	sum := make([]int, l+1)
	for i := 1; i <= l; i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}
	type pair struct {
		i int
		m int
	}
	dp := map[pair]int{}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i == l {
			return 0
		}
		if v, ok := dp[pair{i, m}]; ok {
			return v
		}
		maxVal := math.MinInt
		for x := 1; x <= 2*m; x++ {
			if i+x > l {
				break
			}
			maxVal = max(maxVal, sum[i+x]-sum[i]-dfs(i+x, max(x, m)))
		}
		dp[pair{i, m}] = maxVal
		return maxVal
	}
	return (sum[l] + dfs(0, 1)) / 2
}
