package math

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	l := len(barcodes)
	m := map[int]int{}
	max := 0
	num := 0
	for _, v := range barcodes {
		m[v]++
		if m[v] > max {
			max = m[v]
			num = v
		}
	}
	if max < l/2 {
		num = -1
	}
	ans := []int{}
	if num != -1 {
		for _, v := range barcodes {
			if v == num {
				continue
			}
			ans = append(ans, num)
			ans = append(ans, v)
		}
		if m[num] > l/2 {
			ans = append(ans, num)
		}
		return ans
	}
	sort.Ints(barcodes)
	for i := 0; i < l/2; i++ {

		ans = append(ans, barcodes[l/2+i])
		ans = append(ans, barcodes[i])
	}
	if l%2 != 0 {
		ans = append(ans, barcodes[l-1])
	}
	return ans
}
