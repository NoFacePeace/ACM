package array

func numTimesAllBlue(flips []int) int {
	n := len(flips)
	if n == 0 {
		return 0
	}
	cnt := 0
	max := flips[0]
	ans := 0
	for i := 0; i < n; i++ {
		cnt++
		if flips[i] > max {
			max = flips[i]
		}
		if max == cnt {
			ans++
		}
	}
	return ans
}
