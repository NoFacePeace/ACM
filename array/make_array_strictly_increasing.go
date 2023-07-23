package array

import "sort"

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	cnt := 0
	j := 0
	for i := 0; i < len(arr1)-1; i++ {
		if arr1[i] < arr1[i+1] {
			for j < len(arr2) {
				if arr2[j] > arr1[i] {
					break
				}
				j++
			}
			continue
		}
		if j == len(arr2) {
			return -1
		}
		if arr2[j] < arr1[i+1] {
			cnt++
			arr1[i] = arr2[j]
			for j < len(arr2) {
				if arr2[j] > arr1[i] {
					break
				}
				j++
			}
			continue
		}
		arr1[i] = arr2[j]
		cnt++
		for j < len(arr2) {
			if arr2[j] > arr1[i] {
				break
			}
			j++
		}
		if j == len(arr2) {
			return -1
		}
		arr1[i+1] = arr2[j]
		cnt++
	}
	return cnt
}
