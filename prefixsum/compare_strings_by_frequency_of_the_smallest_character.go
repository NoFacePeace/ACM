package prefixsum

func numSmallerByFrequency(queries []string, words []string) []int {
	f := func(s string) int {
		n := len(s)
		if n == 0 {
			return 0
		}
		c := s[0]
		cnt := 1
		for i := 1; i < n; i++ {
			if s[i] == c {
				cnt++
				continue
			}
			if s[i] > c {
				continue
			}
			c = s[i]
			cnt = 1
		}
		return cnt
	}
	arr := make([]int, 12)
	for _, s := range words {
		cnt := f(s)
		arr[cnt]++
	}
	for i := 9; i >= 0; i-- {
		arr[i] += arr[i+1]
	}

	ans := []int{}
	for _, s := range queries {
		cnt := f(s)
		ans = append(ans, arr[cnt+1])
	}
	return ans
}
