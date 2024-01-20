package math

func tupleSameProduct(nums []int) int {
	m := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			val := nums[i] * nums[j]
			m[val]++
		}
	}
	ans := 0
	f := func(num int) int {
		return num * (num - 1) / 2 * 8
	}
	for _, v := range m {
		ans += f(v)
	}
	return ans
}
