package array

func rob2(nums []int) int {
	l := len(nums)
	tmp := []int{}
	tmp = append(tmp, nums...)
	for i := 0; i < l-1; i++ {
		max := 0
		for j := 0; j < i-1; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		nums[i] += max
	}
	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}
	tmp = tmp[:l-1]
	for i := 0; i < l-1; i++ {
		max := 0
		for j := 0; j < i-1; j++ {
			if tmp[j] > max {
				max = tmp[j]
			}
		}
		tmp[i] += max
	}
	if nums[l-2] > tmp[l-2] {
		return nums[l-2]
	}
	return tmp[l-2]
}
