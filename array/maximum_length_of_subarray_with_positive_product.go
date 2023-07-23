package array

func getMaxLen(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	negative := 0
	positive := 0
	last := 1
	ans := 0
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			positive = 0
			negative = 0
			last = 1
			continue
		}
		if last > 0 {
			if nums[i] > 0 {
				positive++
				if negative != 0 {
					negative++
				}
			} else {
				n := negative
				p := positive
				negative = p + 1
				if n != 0 {
					positive = n + 1
				} else {
					positive = 0
				}
			}
		} else {
			if nums[i] > 0 {
				negative = negative + 1
				if positive != 0 {
					positive++
				} else {
					positive = 1
				}
			} else {
				p := positive
				n := negative
				positive = n + 1
				if p != 0 {
					negative = p + 1
				} else {
					negative = 1
				}
			}
		}
		if positive > ans {
			ans = positive
		}
	}
	return ans
}
