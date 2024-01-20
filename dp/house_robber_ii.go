package dp

func robii(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	f := func(arr []int) int {
		n := len(arr)
		if n == 0 {
			return 0
		}
		if n == 1 {
			return arr[0]
		}
		if n == 2 {
			return max(arr[0], arr[1])
		}
		first := arr[0]
		second := max(arr[0], arr[1])
		for i := 2; i < n; i++ {
			first, second = second, max(first+arr[i], second)
		}
		return second
	}
	return max(f(nums[1:]), f(nums[:len(nums)-1]))
}
