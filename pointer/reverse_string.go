package pointer

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		tmp := s[l]
		s[l] = s[r]
		s[r] = tmp
		l++
		r--
	}
}
