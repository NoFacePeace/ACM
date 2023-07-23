package array

func numberOfArithmeticSlices(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return 0
	}
	dist := nums[1] - nums[0]
	arr := [][]int{}
	cnt := 2
	a := []int{nums[0], nums[1]}
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == dist {
			cnt++
			a = append(a, nums[i])
			continue
		}
		if cnt >= 3 {
			arr = append(arr, a)
		}
		dist = nums[i] - nums[i-1]
		a = []int{nums[i], nums[i-1]}
		cnt = 2
	}
	if len(a) >= 3 {
		arr = append(arr, a)
	}
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := 3; j <= len(arr[i]); j++ {
			ans += len(arr[i]) - j + 1
		}
	}
	return ans
}
