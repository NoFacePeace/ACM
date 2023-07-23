package hash

func findMaxK(nums []int) int {
	ans := -1
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	m := map[int]int{}
	for _, v := range nums {
		if m[abs(v)] == v {
			continue
		}
		m[abs(v)] += v
		if m[abs(v)] != 0 {
			continue
		}
		if abs(v) > ans {
			ans = abs(v)
		}
	}
	return ans
}
