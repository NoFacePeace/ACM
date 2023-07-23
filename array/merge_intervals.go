package array

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	if len(intervals) == 1 {
		return [][]int{
			intervals[0],
		}
	}
	intervals = arraySort(intervals)
	l := intervals[0][0]
	r := intervals[0][1]
	ans := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] > r {
			ans = append(ans, []int{l, r})
			l = intervals[i][0]
			r = intervals[i][1]
			continue
		}
		if intervals[i][1] > r {
			r = intervals[i][1]
		}
	}
	ans = append(ans, []int{l, r})
	return ans
}

func arraySort(arr [][]int) [][]int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr [][]int, left, right int) [][]int {
	if left >= right {
		return arr
	}
	arr, p := partition(arr, left, right)
	arr = quickSort(arr, left, p-1)
	arr = quickSort(arr, p+1, right)
	return arr
}
func partition(arr [][]int, left, right int) ([][]int, int) {
	pivot := arr[right]
	i := 0
	for j := 0; j < right; j++ {
		if arr[j][0] > pivot[0] {
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
	}
	arr[right], arr[i] = arr[i], pivot
	return arr, i
}
