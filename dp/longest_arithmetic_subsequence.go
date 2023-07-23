package dp

func longestArithSeqLength(nums []int) int {
	l := len(nums)
	dp := make([][]int, l)
	max := 1
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
		for j := 0; j < i; j++ {
			dist := nums[i] - nums[j]
			dp[j][i] = 1
			for k := j - 1; k >= 0; k-- {
				if nums[j]-nums[k] == dist {
					dp[j][i] += dp[k][j]
					break
				}
			}
			if dp[j][i] > max {
				max = dp[j][i]
			}
		}
	}
	return max + 1
}
