package dp

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	for _, v := range nums {
		if v > target {
			continue
		}
		dp[v] = 1
	}
	for i := 0; i <= target; i++ {
		for _, v := range nums {
			if i-v < 0 {
				continue
			}
			dp[i] += dp[i-v]
		}
	}
	return dp[target]
}
