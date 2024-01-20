package bitwise

func singleNumbers(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	val := nums[0]
	for i := 1; i < n; i++ {
		val ^= nums[i]
	}
	i := 1
	for val&i == 0 {
		i <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		if i&v == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}
