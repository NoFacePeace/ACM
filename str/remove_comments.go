package str

import "strings"

func removeComments(source []string) []string {
	src := strings.Join(source, `\n`)
	dst := ""
	n := len(src)
	i := 0
	isLine := func(i int) bool {
		if i+1 >= n {
			return false
		}
		if src[i] == '/' && src[i+1] == '/' {
			return true
		}
		return false
	}
	isBlock := func(i int) bool {
		if i+1 >= n {
			return false
		}
		if src[i] == '/' && src[i+1] == '*' {
			return true
		}
		return false
	}
	for i < n {
		if isLine(i) {
			i += 2
			for i < n {
				if i+1 >= n {
					i++
					break
				}
				if src[i:i+2] == `\n` {
					break
				}
				i++
			}
			continue
		}
		if isBlock(i) {
			i += 2
			for i < n && i+1 < n {
				if src[i] == '*' && src[i+1] == '/' {
					i += 2
					break
				}
				i++
			}
			continue
		}
		dst += string(src[i])
		i++
	}
	tmp := strings.Split(dst, `\n`)
	ans := []string{}
	for _, s := range tmp {
		if s == "" {
			continue
		}
		ans = append(ans, s)
	}
	return ans
}
