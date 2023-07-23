package math

import "sort"

func nthUglyNumber(n int) int {
	q := []int{1}
	i := 0
	m := map[int]bool{}
	var num int
	for i < n {
		i++
		num = q[0]
		q = q[1:]
		if !m[num*2] {
			q = append(q, num*2)
			m[num*2] = true
		}
		if !m[num*3] {
			q = append(q, num*3)
			m[num*3] = true
		}
		if !m[num*5] {
			q = append(q, num*5)
			m[num*5] = true
		}
		sort.Ints(q)
	}
	return num
}
