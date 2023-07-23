package array

func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j >= l {
				break
			}
			if dp[i+j] == 0 {
				dp[i+j] = dp[i] + 1
				continue
			}
			if dp[i+j] > dp[i]+1 {
				dp[i+j] = dp[i] + 1
			}
		}
	}
	return dp[l-1]
}
