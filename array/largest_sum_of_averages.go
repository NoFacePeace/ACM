package array

import "math"

func largestSumOfAverages(nums []int, k int) float64 {
	l := len(nums)
	prefix := make([]float64, l+1)
	for i := 0; i < l; i++ {
		prefix[i+1] = prefix[i] + float64(nums[i])
	}
	dp := make([]float64, l+1)
	for i := 1; i <= l; i++ {
		dp[i] = prefix[i] / float64(i)
	}
	for j := 2; j <= k; j++ {
		for i := l; i >= j; i-- {
			for x := j - 1; x < i; x++ {
				dp[i] = math.Max(dp[i], dp[x]+(prefix[i]-prefix[x])/float64(i-x))
			}
		}
	}
	return dp[l]
}
