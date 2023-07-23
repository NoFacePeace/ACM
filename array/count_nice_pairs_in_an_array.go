package array

func countNicePairs(nums []int) int {
	rev := func(num int) int {
		n := 0
		for num != 0 {
			n *= 10
			n += num % 10
			num /= 10
		}
		return n
	}
	m := map[int]int{}
	for _, v := range nums {
		m[v-rev(v)]++
	}
	mod := int(1e9 + 7)
	ans := 0
	for _, v := range m {
		if v == 1 {
			continue
		}
		ans += (v - 1) * v / 2
	}
	return ans % mod
}
