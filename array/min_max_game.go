package array

func minMaxGame(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	for len(nums) > 1 {
		tmp := []int{}
		for i := 0; i < len(nums); i += 2 {
			if i/2%2 == 0 {
				tmp = append(tmp, min(nums[i], nums[i+1]))
			} else {
				tmp = append(tmp, max(nums[i], nums[i+1]))
			}
		}
		nums = tmp
	}
	return nums[0]
}
