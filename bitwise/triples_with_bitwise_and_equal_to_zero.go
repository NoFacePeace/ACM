package bitwise

func countTriplets(nums []int) int {
	l := len(nums)
	ans := 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			ans += 1
		}
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]&nums[j] == 0 {
				ans += 6
			}
		}
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i]&nums[j]&nums[k] == 0 {
					ans += 6
				}
			}
		}
	}
	return ans
}
