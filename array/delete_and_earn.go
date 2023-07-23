package array

import "sort"

func deleteAndEarn(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v] += v
	}
	arr := [][2]int{}
	for k, v := range m {
		arr = append(arr, [2]int{k, v})
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a][0] < arr[b][0]
	})
	for i := 0; i < len(arr); i++ {
		max := 0
		for j := 0; j < i; j++ {
			if arr[i][0]-1 == arr[j][0] {
				continue
			}
			if arr[j][1] > max {
				max = arr[j][1]
			}
		}
		arr[i][1] += max
	}
	max := 0
	for i := 0; i < len(arr); i++ {
		if arr[i][1] > max {
			max = arr[i][1]
		}
	}
	return max
}
