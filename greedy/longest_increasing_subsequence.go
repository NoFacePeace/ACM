package greedy

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort := []int{}
	sort = append(sort, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > sort[len(sort)-1] {
			sort = append(sort, nums[i])
			continue
		}
		p := binarySearch(sort, nums[i])
		sort[p] = nums[i]
	}
	return len(sort)
}

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	p := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			p = mid
			l = mid + 1
			continue
		}
		r = mid - 1
	}
	return p + 1
}
