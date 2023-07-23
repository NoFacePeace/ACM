package slidingwindow

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	m := map[byte]int{}
	m[s[0]]++
	l := 0
	r := 0
	max := 1
	satisfy := func() bool {
		cnt := 0
		max := 0
		for _, v := range m {
			cnt += v
			if v > max {
				max = v
			}
		}
		return cnt-max <= k
	}
	for r < len(s) {
		if satisfy() {
			if r-l+1 > max {
				max = r - l + 1
			}
			r++
			if r == len(s) {
				break
			}
			m[s[r]]++
		} else {
			m[s[l]]--
			l++
		}
	}
	return max
}
