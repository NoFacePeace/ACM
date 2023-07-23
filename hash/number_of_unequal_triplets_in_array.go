package hash

func unequalTriplets(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	ans := 0
	n := len(nums)
	l := 0
	for _, v := range m {
		ans += l * v * (n - l - v)
		l += v
	}
	return ans
}
