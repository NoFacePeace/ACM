package array

func trap(height []int) int {
	l := len(height)
	if l == 0 {
		return 0
	}
	var f func(arr []int, status int) int
	f = func(arr []int, status int) int {
		l := len(arr)
		if l == 0 {
			return 0
		}
		max := 0
		for i := 0; i < l; i++ {
			if arr[i] > arr[max] {
				max = i
			}
		}
		sum := 0
		if status == 0 {
			sum += f(arr[:max], 0)
			for i := max; i < l; i++ {
				sum += arr[max] - arr[i]
			}
		} else {
			sum += f(arr[max+1:], 1)
			for i := 0; i < max; i++ {
				sum += arr[max] - arr[i]
			}
		}
		return sum
	}
	max := 0
	ans := 0
	for i := 0; i < l; i++ {
		if height[i] > height[max] {
			max = i
		}
	}
	ans += f(height[:max], 0)
	ans += f(height[max+1:], 1)
	return ans
}
