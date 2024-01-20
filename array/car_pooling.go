package array

func carPooling(trips [][]int, capacity int) bool {
	from := map[int]int{}
	to := map[int]int{}
	for _, v := range trips {
		n, f, t := v[0], v[1], v[2]
		from[f] += n
		to[t] += n
	}
	c := 0
	for i := 0; i <= 1000; i++ {
		c -= to[i]
		c += from[i]
		if c > capacity {
			return false
		}
	}
	return true
}
