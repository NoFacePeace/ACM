package math

func minSubarray(nums []int, p int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	mod := sum % p
	if mod == 0 {
		return 0
	}
	sum = 0
	m := map[int]int{0: -1}
	min := len(nums)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		val := (sum - mod + p) % p
		if _, ok := m[val]; ok {
			if i-m[val] < min {
				min = i - m[val]
			}
		}
		m[(sum+p)%p] = i
	}
	if min == len(nums) {
		return -1
	}
	return min
}
