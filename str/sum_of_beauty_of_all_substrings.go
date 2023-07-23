package str

import "math"

func beautySum(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		m := map[byte]int{}
		for j := i; j < len(s); j++ {
			max := math.MinInt
			min := math.MaxInt
			m[s[j]]++
			if len(m) == 1 {
				continue
			}
			for _, v := range m {
				if v > max {
					max = v
				}
				if v < min {
					min = v
				}
			}
			ans += max - min
		}
	}
	return ans
}
