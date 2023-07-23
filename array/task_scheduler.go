package array

import "sort"

func leastInterval(tasks []byte, n int) int {
	m := map[byte]int{}
	for _, v := range tasks {
		m[v]++
	}
	arr := []int{}
	for _, v := range m {
		arr = append(arr, v)
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a] > arr[b]
	})
	ans := 0
	for len(arr) > 0 {
		i := 0
		for i <= n {
			if i >= len(arr) {
				break
			}
			arr[i]--
			i++
			ans++
		}
		tmp := arr
		arr = nil
		for i := 0; i < len(tmp); i++ {
			if tmp[i] != 0 {
				arr = append(arr, tmp[i])
			}
		}
		sort.Slice(arr, func(a, b int) bool {
			return arr[a] > arr[b]
		})
		if len(arr) > 0 {
			ans += n + 1 - i
		}
	}
	return ans
}
