package str

func letterCombinations(digits string) []string {
	ans := []string{}
	m := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	var bk func(digits string, index int, s string)
	bk = func(digits string, index int, s string) {
		if index >= len(digits) {
			ans = append(ans, s)
			return
		}
		for i := 0; i < len(m[digits[index]]); i++ {
			bk(digits, index+1, s+string(m[digits[index]][i]))
		}
	}
	bk(digits, 0, "")
	return ans
}
