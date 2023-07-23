package array

func findKthLargest(nums []int, k int) int {
	partition := func(nums []int, left, right int) ([]int, int) {
		pivot := nums[right]
		i := left
		for j := left; j < right; j++ {
			if nums[j] < pivot {
				nums[j], nums[i] = nums[i], nums[j]
				i++
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
		return nums, i
	}
	l := 0
	r := len(nums) - 1
	k = len(nums) - k
	for l <= r {
		nums, p := partition(nums, l, r)
		if p == k {
			return nums[p]
		}
		if p > k {
			r = p - 1
			continue
		}
		l = p + 1
	}
	return -1
}
