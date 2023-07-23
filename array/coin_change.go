package array

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		min := math.MaxInt - 1
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			if dp[i-coins[j]] < min {
				min = dp[i-coins[j]]
			}
		}
		dp[i] = min + 1
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
