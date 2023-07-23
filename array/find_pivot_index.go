package array

func pivotIndex(nums []int) int {
	l := len(nums)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			arr[i] = nums[i]
			continue
		}
		arr[i] = arr[i-1] + nums[i]
	}
	for i := 0; i < l; i++ {
		if i == 0 {
			if arr[l-1]-arr[i] == 0 {
				return i
			}
			continue
		}
		if arr[i-1] == arr[l-1]-arr[i] {
			return i
		}
	}
	return -1
}
