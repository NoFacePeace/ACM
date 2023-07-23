package str

func digitCount(num string) bool {
	m := map[byte]int{}
	for i := 0; i < len(num); i++ {
		m[num[i]] += 1
	}
	for i := 0; i < len(num); i++ {
		c := byte('0' + m[byte('0'+i)])
		if c != num[i] {
			return false
		}
	}
	return true
}
