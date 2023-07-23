package greedy

func wiggleMaxLength(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	tmp := nums
	nums = []int{tmp[0]}
	for i := 1; i < len(tmp); i++ {
		if tmp[i] == tmp[i-1] {
			continue
		}
		nums = append(nums, tmp[i])
	}
	l = len(nums)
	up := 0
	down := 0
	for i := 1; i < l; i++ {
		if i == l-1 && nums[i] > nums[i-1] {
			up++
			continue
		}
		if i == l-1 && nums[i] < nums[i-1] {
			down++
			continue
		}
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			up++
			continue
		}
		if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			down++
			continue
		}
	}
	return up + down + 1
}
