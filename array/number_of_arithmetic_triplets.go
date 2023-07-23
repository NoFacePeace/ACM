package array

func arithmeticTriplets(nums []int, diff int) int {
	ans := 0
	l := len(nums)
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			if nums[j]-nums[i] < diff {
				continue
			}
			if nums[j]-nums[i] > diff {
				break
			}
			for k := j + 1; k < l; k++ {
				if nums[k]-nums[j] == diff {
					ans++
					break
				}
				if nums[k]-nums[j] > diff {
					break
				}
			}
		}
	}
	return ans
}
