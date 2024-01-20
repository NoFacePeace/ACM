package bitwise

func subtractProductAndSum(n int) int {
	mul := 1
	sum := 0
	for n != 0 {
		bit := n % 10
		n /= 10
		mul *= bit
		sum += bit
	}
	return mul - sum
}
