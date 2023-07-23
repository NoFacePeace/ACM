package array

func reachNumber(target int) int {
	if target == 0 {
		return 0
	}
	q := []int{}
	i := 0
	q = append(q, 0)
	for len(q) > 0 {
		i++
		tmp := q
		q = nil
		for j := 0; j < len(tmp); j++ {
			num := tmp[j] + i
			if num == target {
				return i
			}
			q = append(q, num)
			num = tmp[j] - i
			if num == target {
				return i
			}
			q = append(q, num)
		}
	}
	return 0
}
