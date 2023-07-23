package hash

func findSubarrays(nums []int) bool {
	ans := false
	m := map[int]bool{}
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if _, ok := m[sum]; ok {
			ans = true
			break
		}
		m[sum] = true
	}
	return ans
}
