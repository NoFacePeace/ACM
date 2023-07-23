package array

func waysToMakeFair(nums []int) int {
	l1 := 0
	l2 := 0
	l3 := 0
	l4 := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			l4 += nums[i]
		} else {
			l3 += nums[i]
		}
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			if l1+l4-nums[i] == l2+l3 {
				ans++
			}
			l2 += nums[i]
			l4 -= nums[i]
		} else {
			if l1+l4 == l2+l3-nums[i] {
				ans++
			}
			l1 += nums[i]
			l3 -= nums[i]
		}
	}
	return ans
}
