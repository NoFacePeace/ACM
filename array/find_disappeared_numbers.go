package array

func findDisappearedNumbers(nums []int) []int {
	var ans []int
	if len(nums) == 0 {
		return ans
	}
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] = nums[v] + n
	}
	for i, v := range nums {
		if v > n {
			continue
		}
		ans = append(ans, i+1)
	}
	return ans
}
