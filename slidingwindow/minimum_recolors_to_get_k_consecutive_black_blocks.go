package slidingwindow

import "math"

func minimumRecolors(blocks string, k int) int {
	if len(blocks) < k {
		return 0
	}
	i := 0
	j := 0
	w := 0
	b := 0
	ans := math.MaxInt
	for j < len(blocks) {
		if blocks[j] == 'W' {
			w++
		} else {
			b++
		}
		if j-i >= k {
			if blocks[i] == 'W' {
				w--
			} else {
				b--
			}
			i++
		}
		if j-i+1 == k && w < ans {
			ans = w
		}
		j++
	}
	return ans
}
