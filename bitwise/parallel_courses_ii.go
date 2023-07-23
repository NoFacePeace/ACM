package bitwise

func minNumberOfSemesters(n int, relations [][]int, k int) int {
	m := map[int]int{}
	for i := 1; i <= n; i++ {
		m[i] = 0
	}
	for _, r := range relations {
		v, e := r[0], r[1]
		bit := 1 << (v - 1)
		m[e] |= bit
	}
	cnt := 0
	q := []int{}
	for len(m) > 0 {
		for e, v := range m {
			if v == 0 {
				q = append(q, e)
			}
		}
		for i := 0; i < k; i++ {
			if len(q) == 0 {
				break
			}
			v := q[0]
			delete(m, v)
			for e := range m {
				if m[e]>>(v-1)&1 == 1 {
					bit := 1 << (v - 1)
					m[e] = m[e] ^ bit
				}
			}
			q = q[1:]
		}
		cnt++
	}
	return cnt
}
