package str

import "strings"

func isCircularSentence(sentence string) bool {
	arr := strings.Split(sentence, " ")
	n := len(arr)
	if n == 0 {
		return true
	}
	if n == 1 {
		s := arr[0]
		return s[0] == s[len(s)-1]
	}
	first := arr[0]
	last := arr[n-1]
	if first[0] != last[len(last)-1] {
		return false
	}
	for i := 1; i < n; i++ {
		s1 := arr[i-1]
		s2 := arr[i]
		if s1[len(s1)-1] != s2[0] {
			return false
		}
	}
	return true
}
