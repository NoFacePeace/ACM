package array

func findDuplicate(nums []int) int {
	l := len(nums)
	for i := 0; i < l; i++ {
		tmp := nums[i] % (l + 1)
		nums[tmp-1] += (l + 1)
	}
	for i := 0; i < l; i++ {
		if nums[i] > 2*(l+1) {
			return i + 1
		}
	}
	return -1
}
