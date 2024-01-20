package dp

import "sort"

func sumOfPower(nums []int) int {
	sort.Ints(nums)
	sum := 0
	ans := 0
	mod := int(1e9) + 7
	for _, v := range nums {
		tmp := sum
		sum += v
		ans += sum * v % mod * v % mod
		sum += tmp
		sum %= mod
	}
	return ans % mod
}
