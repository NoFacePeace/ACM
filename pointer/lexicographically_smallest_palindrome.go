package pointer

func makeSmallestPalindrome(s string) string {
	n := len(s)
	s1 := ""
	s2 := ""
	for i := 0; i < n/2; i++ {
		if s[i] <= s[n-i-1] {
			s1 += s[i : i+1]
			s2 = s[i:i+1] + s2
		} else {
			s1 += s[n-i-1 : n-i]
			s2 = s[n-i-1:n-i] + s2
		}
	}
	if n%2 != 0 {
		s1 += s[n/2 : n/2+1]
	}
	return s1 + s2
}
