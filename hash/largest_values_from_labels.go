package hash

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	type pair struct {
		value int
		label int
	}
	arr := []pair{}
	n := len(values)
	for i := 0; i < n; i++ {
		p := pair{
			value: values[i],
			label: labels[i],
		}
		arr = append(arr, p)
	}
	sort.Slice(arr, func(a, b int) bool {
		return arr[a].value < arr[b].value
	})
	m := map[int]int{}
	ans := 0
	cnt := 0
	for i := n - 1; i >= 0; i-- {
		if cnt >= numWanted {
			break
		}
		value := arr[i].value
		label := arr[i].label
		if m[label] >= useLimit {
			continue
		}
		ans += value
		m[label]++
		cnt++
	}
	return ans
}
