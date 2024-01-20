package hash

func maximumSum(nums []int) int {
	m := map[int]int{}
	ans := -1
	sum := func(num int) int {
		val := 0
		for num != 0 {
			val += num % 10
			num /= 10
		}
		return val
	}
	for _, v := range nums {
		s := sum(v)
		if val, ok := m[s]; ok {
			if val+v > ans {
				ans = val + v
			}
			if v > val {
				m[s] = v
			}
		} else {
			m[s] = v
		}
	}
	return ans
}
