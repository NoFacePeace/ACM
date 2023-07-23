package hash

func mostFrequentEven(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if v%2 != 0 {
			continue
		}
		m[v]++
	}
	max := 0
	ans := -1
	for k, v := range m {
		if v < max {
			continue
		}
		if v > max {
			max = v
			ans = k
			continue
		}
		if k < ans {
			ans = k
		}
	}
	return ans
}
