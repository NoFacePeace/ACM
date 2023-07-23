package array

func missingNumber(nums []int) int {
	sum1 := 0
	sum2 := 0
	for i := 0; i < len(nums); i++ {
		sum1 += nums[i]
		sum2 += i + 1
	}
	return sum2 - sum1
}
