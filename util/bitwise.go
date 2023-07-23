package util

func numberOfLeadingZeros(num int) int {
	if num == 0 {
		return 32
	}
	n := 1
	if num>>16 == 0 {
		n += 16
		num = num << 16
	}
	if num>>24 == 0 {
		n += 8
		num = num << 8
	}
	if num>>28 == 0 {
		n += 4
		num = num << 4
	}
	if num>>30 == 0 {
		n += 2
		num = num << 2
	}
	n -= num >> 31
	return n
}
