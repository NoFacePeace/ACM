package array

import "sort"

func sortPeople(names []string, heights []int) []string {
	type person struct {
		name   string
		height int
	}
	arr := []person{}
	for i := 0; i < len(names); i++ {
		p := person{
			name:   names[i],
			height: heights[i],
		}
		arr = append(arr, p)
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a].height > arr[b].height
	})
	ans := []string{}
	for _, v := range arr {
		ans = append(ans, v.name)
	}
	return ans
}
