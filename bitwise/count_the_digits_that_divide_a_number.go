package bitwise

func countDigits(num int) int {
	cnt := 0
	n := num
	for n != 0 {
		mod := n % 10
		n /= 10
		if num%mod == 0 {
			cnt++
		}
	}
	return cnt
}
