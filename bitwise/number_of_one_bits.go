package bitwise

func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		if num&1 != 0 {
			ans += 1
		}
		num >>= 1
	}
	return ans
}
