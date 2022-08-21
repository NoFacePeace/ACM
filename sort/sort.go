package sort

func quickSort(arr []int, left, right int) []int {
	if left >= right {
		return arr
	}
	arr, mid := partition(arr, left, right)
	quickSort(arr, left, mid-1)
	quickSort(arr, mid+1, right)
	return arr
}

func partition(arr []int, left, right int) ([]int, int) {
	pivot := arr[right]
	i := left
	for j := i; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}

func binarySearch(arr []int, target int) int {
	if arr == nil {
		return -1
	}
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return -1
}
