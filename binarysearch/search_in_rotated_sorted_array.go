package binarysearch

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target && nums[l] <= target {
			r = mid - 1
			continue
		}
		if nums[mid] > target && nums[l] > target {
			l = mid + 1
			continue
		}
		if nums[mid] < target && nums[r] >= target {
			l = mid + 1
			continue
		}
		r = mid - 1
	}
	return -1
}
