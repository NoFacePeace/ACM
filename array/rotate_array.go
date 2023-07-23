package array

func rotate_array(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	l := len(nums) - 1
	for i := 0; i <= l/2; i++ {
		nums[i], nums[l-i] = nums[l-i], nums[i]
	}
	mid := k - 1
	for i := 0; i <= mid/2; i++ {
		nums[i], nums[mid-i] = nums[mid-i], nums[i]
	}
	for i := 0; i+k <= (k+l)/2; i++ {
		nums[i+k], nums[l-i] = nums[l-i], nums[i+k]
	}
}
