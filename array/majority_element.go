package array

func majorityElement(nums []int) int {
	cnt := 0
	ans := 0
	for _, v := range nums {
		if cnt == 0 {
			ans = v
			cnt = 1
			continue
		}
		if v == ans {
			cnt++
			continue
		}
		cnt--
	}
	return ans
}
