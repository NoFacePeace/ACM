package str

import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
	ans := []string{}
	for _, w := range words {
		arr := strings.Split(w, string(separator))
		for _, v := range arr {
			if v != "" {
				ans = append(ans, v)
			}
		}
	}
	return ans
}
