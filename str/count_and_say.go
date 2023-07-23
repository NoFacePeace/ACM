package str

import "strconv"

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	i := 0
	ans := ""
	count := 1
	for i < len(s) {
		if i+1 == len(s) {
			ans += strconv.Itoa(count) + string(s[i])
			i++
			continue
		}
		if s[i] == s[i+1] {
			count++
			i++
			continue
		}
		ans += strconv.Itoa(count) + string(s[i])
		count = 1
		i++
	}
	return ans
}
