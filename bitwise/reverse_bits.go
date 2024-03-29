package bitwise

func reverseBits(num uint32) uint32 {
	ans := uint32(0)
	for i := 0; i < 32; i++ {
		ans |= (num & 1) << (31 - i)
		num = num >> 1
	}
	return ans
}
