package array

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	m := map[int]bool{}
	for k := range nums {
		if _, ok := m[nums[k]]; ok {
			return true
		}
		m[nums[k]] = true
	}
	return false
}
