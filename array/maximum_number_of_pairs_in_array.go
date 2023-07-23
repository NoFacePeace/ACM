package array

func numberOfPairs(nums []int) []int {
	ans := []int{0, 0}
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	for _, v := range m {
		ans[0] += v / 2
		ans[1] += v % 2
	}
	return ans
}
