package dp

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n)
	}
	sum := 0
	for i := 0; i < n; i++ {
		if i < k-1 {
			sum += nums[i]
			continue
		}
		if i == k-1 {
			sum += nums[i]
			dp[0][i] = sum
			continue
		}
		sum = sum - nums[i-k] + nums[i]
		dp[0][i] = max(dp[0][i-1], sum)
	}
	sum = 0
	for i := k; i < n; i++ {
		if i < 2*k-1 {
			sum += nums[i]
			continue
		}
		if i == 2*k-1 {
			sum += nums[i]
			dp[1][i] = dp[0][i-k] + sum
			continue
		}
		sum = sum - nums[i-k] + nums[i]
		dp[1][i] = max(dp[1][i-1], sum+dp[0][i-k])
	}
	sum = 0
	for i := 2 * k; i < n; i++ {
		if i < 3*k-1 {
			sum += nums[i]
			continue
		}
		if i == 3*k-1 {
			sum += nums[i]
			dp[2][i] = dp[1][i-k] + sum
			continue
		}
		sum = sum - nums[i-k] + nums[i]
		dp[2][i] = max(dp[2][i-1], sum+dp[1][i-k])
	}
	third := n - 1
	for i := n - 1; i > 0; i-- {
		if dp[2][i] != dp[2][i-1] {
			break
		}
		third = i - 1
	}
	second := third - k
	for i := third - k; i > 0; i-- {
		if dp[1][i] != dp[1][i-1] {
			break
		}
		second = i - 1
	}
	first := second - k
	for i := second - k; i > 0; i-- {
		if dp[0][i] != dp[0][i-1] {
			break
		}
		first = i - 1
	}
	return []int{first - k + 1, second - k + 1, third - k + 1}
}
