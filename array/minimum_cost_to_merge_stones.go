package array

import "math"

func mergeStones(stones []int, k int) int {
	l := len(stones)
	if (l-1)%(k-1) != 0 {
		return -1
	}
	prefix := make([]int, l+1)
	for i := 0; i < l; i++ {
		prefix[i+1] = prefix[i] + stones[i]
	}
	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		if left == right {
			return 0
		}
		if dp[left][right] != 0 {
			return dp[left][right]
		}
		min := math.MaxInt
		for i := left; i < right; i += k - 1 {
			val := dfs(left, i) + dfs(i+1, right)
			if val < min {
				min = val
			}
		}
		if (right-left)%(k-1) == 0 {
			min += prefix[right+1] - prefix[left]
		}
		dp[left][right] = min
		return min
	}
	return dfs(0, l-1)
}
