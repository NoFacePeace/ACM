package pointer

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	var ans []byte
	var l int
	if l1 > l2 {
		ans = make([]byte, l1+1)
		l = l1
	} else {
		ans = make([]byte, l2+1)
		l = l2
	}
	i := l1 - 1
	j := l2 - 1
	k := l
	carry := byte(0)
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			ans[k] = num1[i] + num2[j] - '0' + carry
			i--
			j--
		} else if i >= 0 {
			ans[k] = num1[i] + carry
			i--
		} else {
			ans[k] = num2[j] + carry
			j--
		}
		if ans[k] > '9' {
			ans[k] -= 10
			carry = 1
		} else {
			carry = 0
		}
		k--
	}
	if carry > byte(0) {
		ans[0] = '1'
	} else {
		ans = ans[1:]
	}
	return string(ans)
}
