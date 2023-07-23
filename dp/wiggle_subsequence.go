package dp

func wiggleMaxLength(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	up := make([]int, l)
	down := make([]int, l)
	up[0] = 1
	down[0] = 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			up[i] = up[i-1]
			down[i] = down[i-1]
			continue
		}
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
			continue
		}
		up[i] = up[i-1]
		down[i] = max(down[i-1], up[i-1]+1)
	}
	return max(up[l-1], down[l-1])
}
