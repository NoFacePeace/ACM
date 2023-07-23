package bitwise

func hammingDistance(x int, y int) int {
	v := x ^ y
	ans := 0
	for v != 0 {
		v = v & (v - 1)
		ans += 1
	}
	return ans
}
