package array

func captureForts(forts []int) int {
	max := 0
	n := len(forts)
	cnt := 0
	p := 0
	for i := 0; i < n; i++ {
		if forts[i] == 0 {
			cnt++
			continue
		}
		if forts[i] == 1 {
			if p == 0 {
				p = forts[i]
				cnt = 0
				continue
			}
			if p == 1 {
				cnt = 0
				continue
			}
			if cnt > max {
				max = cnt
			}
			p = forts[i]
			cnt = 0
			continue
		}
		if p == 0 {
			p = forts[i]
			cnt = 0
			continue
		}
		if p == -1 {
			cnt = 0
			continue
		}
		if cnt > max {
			max = cnt
		}
		p = forts[i]
		cnt = 0
	}
	return max
}
