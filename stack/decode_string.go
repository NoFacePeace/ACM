package stack

import "strconv"

func decodeString(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
			continue
		}
		ss := ""
		for len(stack) > 0 {
			l := len(stack)
			c := stack[l-1]
			stack = stack[:l-1]
			if c == '[' {
				break
			}
			ss = string(c) + ss
		}
		sss := ""
		for len(stack) > 0 {
			l := len(stack)
			c := stack[l-1]
			if c <= '9' && c >= '0' {
				sss = string(c) + sss
				stack = stack[:l-1]
				continue
			}
			break
		}
		num, _ := strconv.Atoi(sss)
		ssss := ""
		for j := 0; j < num; j++ {
			ssss += ss
		}
		for j := 0; j < len(ssss); j++ {
			stack = append(stack, ssss[j])
		}
	}
	ans := ""
	for _, v := range stack {
		if v <= 'z' && v >= 'a' {
			ans += string(v)
		}
	}
	return ans
}
