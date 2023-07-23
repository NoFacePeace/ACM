package sort

// QuickSort quick sort support generics
func QuickSort[T any](arr []T, f func(a, b T) bool) []T {
	inArr := []T{}
	inArr = append(inArr, arr...)
	if len(inArr) <= 1 {
		return inArr
	}
	inArr = qSort(inArr, 0, len(inArr)-1, f)
	return inArr
}

// MergeSort merge sort support generics
func MergeSort[T any](arr []T, f func(a, b T) bool) []T {
	inArr := []T{}
	inArr = append(inArr, arr...)
	if len(inArr) <= 1 {
		return inArr
	}
	inArr = mSort(inArr, f)
	return inArr
}

func mSort[T any](arr []T, f func(a, b T) bool) []T {
	if len(arr) <= 1 {
		return arr
	}
	l := len(arr)
	left := mSort(arr[:l/2], f)
	right := mSort(arr[l/2:], f)
	return merge(left, right, f)
}

func BinarySearch[T any](arr []T, target T, equal func(a, b T) int) int {
	if len(arr) <= 0 {
		return -1
	}
	l := 0
	r := len(arr)
	for l <= r {
		mid := (l + r) / 2
		if equal(arr[mid], target) == 0 {
			return mid
		}
		if equal(arr[mid], target) == 1 {
			r = mid - 1
			continue
		}
		l = mid + 1
	}
	return -1
}
func merge[T any](arr1, arr2 []T, f func(a, b T) bool) []T {
	arr := []T{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if f(arr1[i], arr2[j]) {
			arr = append(arr, arr2[j])
			j++
			continue
		}
		arr = append(arr, arr1[i])
		i++
	}
	for i < len(arr1) {
		arr = append(arr, arr1[i])
		i++
	}
	for j < len(arr2) {
		arr = append(arr, arr2[j])
		j++
	}
	return arr
}

func qSort[T any](arr []T, left, right int, f func(a, b T) bool) []T {
	if left >= right {
		return arr
	}
	arr, p := partition(arr, left, right, f)
	arr = qSort(arr, left, p-1, f)
	arr = qSort(arr, p+1, right, f)
	return arr
}

func partition[T any](arr []T, left int, right int, f func(a, b T) bool) ([]T, int) {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if f(arr[j], pivot) {
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
	}
	arr[i], arr[right] = arr[right], arr[i]
	return arr, i
}
