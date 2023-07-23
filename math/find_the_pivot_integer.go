package math

func pivotInteger(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	tmp := 0
	for i := 1; i <= n; i++ {
		tmp += i
		if tmp == sum-tmp+i {
			return i
		}
	}
	return -1
}
