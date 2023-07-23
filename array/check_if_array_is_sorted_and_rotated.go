package array

func check(nums []int) bool {
	i := 0
	for i < len(nums)-1 {
		if nums[i] > nums[i+1] {
			break
		}
		i++
	}
	if i == len(nums)-1 {
		return true
	}
	left := nums[:i+1]
	right := nums[i+1:]
	nums = append(right, left...)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		return false
	}
	return true
}
