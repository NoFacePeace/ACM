package str

import "strings"

func maskPII(s string) string {
	maskEmail := func(s string) string {
		ans := ""
		for i := 0; i < len(s); i++ {
			c := s[i]
			if c >= 'A' && c <= 'Z' {
				c = 'a' + c - 'A'
			}
			ans += string(c)
		}
		arr := strings.Split(ans, "@")
		if len(arr) != 2 {
			return ""
		}
		return string(arr[0][0]) + "*****" + string(arr[0][len(arr[0])-1]) + "@" + arr[1]
	}
	maskPhone := func(s string) string {
		ans := ""
		cnt := 0
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] >= '0' && s[i] <= '9' {
				cnt++
				if cnt <= 4 {
					ans = string(s[i]) + ans
				}
			}
		}
		cnt -= 10
		tmp := "***-***-"
		if cnt == 0 {
			return tmp + ans
		}
		if cnt == 1 {
			return "+*-" + tmp + ans
		}
		if cnt == 2 {
			return "+**-" + tmp + ans
		}
		return "+***-" + tmp + ans
	}
	arr := strings.Split(s, "@")
	if len(arr) != 2 {
		return maskPhone(s)
	}
	return maskEmail(s)
}
