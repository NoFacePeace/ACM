package array

func findPeakElement(nums []int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			r = mid
			continue
		}
		l = mid + 1
	}
	return l
}
