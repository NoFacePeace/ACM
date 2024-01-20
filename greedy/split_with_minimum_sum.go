package greedy

import "sort"

func splitNum(num int) int {
	arr := []int{}
	for num != 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	sort.Ints(arr)
	num1 := 0
	num2 := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			num1 *= 10
			num1 += arr[i]
		} else {
			num2 *= 10
			num2 += arr[i]
		}
	}
	return num1 + num2
}
