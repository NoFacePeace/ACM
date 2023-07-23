package array

import "math"

func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	max := 1
	ans := math.MaxInt
	zero := 0
	for i := 1; i <= max; i++ {
		cnt := i
		for j := 0; j < n; j++ {
			if vat[j] > max {
				max = vat[j]
			}
			up := 0
			if vat[j] == 0 {
				zero++
			}
			for (bucket[j]+up)*i < vat[j] {
				up++
				cnt++
			}
		}
		if cnt < ans {
			ans = cnt
		}
	}
	if zero == n {
		return 0
	}
	return ans
}
