package array

func sortColors(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
	for j := i; j < len(nums); j++ {
		if nums[j] != 1 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
	for j := i; j < len(nums); j++ {
		if nums[j] != 1 {
			continue
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
	}
}
