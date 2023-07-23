package greedy

func lemonadeChange(bills []int) bool {
	m := map[int]int{
		5:  0,
		10: 0,
	}
	for _, v := range bills {
		if v == 5 {
			m[5]++
			continue
		}
		if v == 10 {
			if m[5] == 0 {
				return false
			}
			m[5]--
			m[10]++
			continue
		}
		if m[10] > 0 && m[5] > 0 {
			m[10]--
			m[5]--
			continue
		}
		if m[5] < 3 {
			return false
		}
		m[5] -= 3
	}
	return true
}
