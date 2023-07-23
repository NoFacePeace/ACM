package str

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 {
		return ""
	}
	if len(num2) == 0 {
		return ""
	}
	if num1 == "0" {
		return "0"
	}
	if num2 == "0" {
		return "0"
	}
	n2 := int(num2[0] - '0')
	bit := 0
	s := ""
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		n := n1*n2 + bit
		bit = n / 10
		c := byte('0' + (n % 10))
		s = string(c) + s
	}
	if bit != 0 {
		c := byte('0' + bit)
		s = string(c) + s
	}
	if len(num2) == 1 {
		return s
	}
	for i := 0; i < len(num2)-1; i++ {
		s += "0"
	}
	add := func(num1, num2 string) string {
		i := len(num1) - 1
		j := len(num2) - 1
		bit := 0
		s := ""
		for i >= 0 && j >= 0 {
			n1 := int(num1[i] - '0')
			n2 := int(num2[j] - '0')
			n := n1 + n2 + bit
			c := byte('0' + (n % 10))
			bit = n / 10
			s = string(c) + s
			i--
			j--
		}
		for i >= 0 {
			n := int(num1[i]-'0') + bit
			c := byte('0' + (n % 10))
			bit = n / 10
			s = string(c) + s
			i--
		}
		for j >= 0 {
			n := int(num2[j]-'0') + bit
			c := byte('0' + (n % 10))
			bit = n / 10
			s = string(c) + s
			j--
		}
		if bit != 0 {
			s = string(byte('0'+bit)) + s
		}
		return s
	}
	return add(s, multiply(num1, num2[1:]))
}
