package str

import "strconv"

func getLucky(s string, k int) int {
	s1 := ""
	for _, c := range s {
		s1 += strconv.Itoa(int(c-'a') + 1)
	}
	num := 0
	for _, c := range s1 {
		num += int(c - '0')
	}
	k--
	for k > 0 {
		tmp := num
		num = 0
		for tmp > 0 {
			num += tmp % 10
			tmp /= 10
		}
		k--
	}
	return num
}
