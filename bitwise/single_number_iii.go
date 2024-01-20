package bitwise

func singleNumberIII(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	bit := xor & (-xor)
	first := 0
	second := 0
	for _, v := range nums {
		if v&bit == 0 {
			first ^= v
		} else {
			second ^= v
		}
	}
	return []int{first, second}
}
