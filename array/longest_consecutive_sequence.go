package array

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]-1]; ok {
			continue
		}
		count := 1
		index := 1
		for {
			if _, ok := m[nums[i]+index]; ok {
				index++
				count++
				continue
			}
			if count > max {
				max = count
			}
			break
		}
	}
	return max
}
