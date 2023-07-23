package array

func averageValue(nums []int) int {
	sum := 0
	cnt := 0
	for _, v := range nums {
		if v%2 != 0 {
			continue
		}
		if v%3 != 0 {
			continue
		}
		sum += v
		cnt++
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
