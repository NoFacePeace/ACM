package bitwise

import "math/bits"

func maximumRows(matrix [][]int, numSelect int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	convert := func(arr []int) int {
		num := 0
		for k, v := range arr {
			if v == 0 {
				continue
			}
			num |= 1 << k
		}
		return num
	}
	arr := []int{}
	for _, v := range matrix {
		num := convert(v)
		arr = append(arr, num)
	}
	max := 0
	limit := 1 << n
	for i := 1; i < limit; i++ {
		if bits.OnesCount(uint(i)) != numSelect {
			continue
		}
		cnt := 0
		for _, v := range arr {
			if i|v == i {
				cnt++
			}
		}
		if cnt > max {
			max = cnt
		}
	}
	return max
}
