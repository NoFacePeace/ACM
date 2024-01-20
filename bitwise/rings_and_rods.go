package bitwise

func countPoints(rings string) int {
	n := len(rings)
	arr := make([]int, 10)
	cnt := 0
	for i := 0; i < n; i += 2 {
		idx := int(rings[i+1] - '0')
		if arr[idx] == 7 {
			continue
		}
		bit := 0
		if rings[i] == 'R' {
			bit = 4
		} else if rings[i] == 'G' {
			bit = 2
		} else {
			bit = 1
		}
		arr[idx] |= bit
		if arr[idx] == 7 {
			cnt++
		}
	}
	return cnt
}
