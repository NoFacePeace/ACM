package dp

func countBits(n int) []int {
	bits := make([]int, n+1)
	hightBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			hightBit = i
		}
		bits[i] = bits[i-hightBit] + 1
	}
	return bits
}
