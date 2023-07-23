package array

func splitArraySameAverage(nums []int) bool {
	var bk func(n1, c1, n2, c2, index int) bool
	l := len(nums)
	bk = func(n1, c1, n2, c2, index int) bool {
		if index == l {
			if c1 == 0 || c2 == 0 {
				return false
			}
			if n1/c1 == n2/c2 {
				return true
			}
			return false
		}
		n1 += nums[index]
		c1 += 1
		if bk(n1, c1, n2, c2, index+1) {
			return true
		}
		n1 -= nums[index]
		c1 -= 1
		n2 += nums[index]
		c2 += 1
		if bk(n1, c1, n2, c2, index+1) {
			return true
		}
		return false
	}
	return bk(0, 0, 0, 0, 0)
}
