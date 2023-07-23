package array

func runningSum(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans = append(ans, nums[i])
			continue
		}
		ans = append(ans, nums[i]+ans[i-1])
	}
	return ans
}
