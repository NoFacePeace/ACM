package array

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < n/2; i++ {
		first := nums[i]
		second := nums[n-i-1]
		str := strconv.Itoa(first) + strconv.Itoa(second)
		num, _ := strconv.Atoi(str)
		ans += num
	}
	if n%2 != 0 {
		ans += nums[n/2]
	}
	return int64(ans)
}
