package hash

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	n := len(nums)
	ans := []int{}
	for i := 0; i < n; i++ {
		val := target - nums[i]
		if _, ok := m[val]; ok {
			ans = append(ans, m[val])
			ans = append(ans, i)
			break
		}
		m[nums[i]] = i
	}
	return ans
}
