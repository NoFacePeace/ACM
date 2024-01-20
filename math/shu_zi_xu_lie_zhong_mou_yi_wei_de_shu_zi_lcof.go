package math

import "strconv"

func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	cnt := 0
	nine := 9
	bit := 1
	for cnt+nine*bit < n {
		cnt += nine * bit
		bit++
		nine *= 10
	}
	n -= cnt
	start := nine / 9
	mod := n % bit
	n /= bit
	start += n
	if mod == 0 {
		start--
		return start % 10
	}
	str := strconv.Itoa(start)
	return int(str[mod-1] - '0')
}
