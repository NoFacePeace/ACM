package other

func majorityElement(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	count := 1
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			ans = nums[i]
			count++
		} else {
			if nums[i] == ans {
				count++
			} else {
				count--
			}
		}
	}
	return ans
}
