package hash

func powerfulIntegers(x int, y int, bound int) []int {
	if bound <= 1 {
		return []int{}
	}
	if x == 1 && y == 1 {
		return []int{2}
	}
	xx := []int{1}
	yy := []int{1}
	i, j := 0, 0
	m := map[int]bool{}
	for xx[i]+yy[j] <= bound {
		m[xx[i]+yy[j]] = true
		if x == 1 {
			j++
			if len(yy) == j {
				yy = append(yy, yy[j-1]*y)
			}
			continue
		}
		if y == 1 {
			i++
			if len(xx) == i {
				xx = append(xx, xx[i-1]*x)
			}
			continue
		}
		j++
		if len(yy) == j {
			yy = append(yy, yy[j-1]*y)
		}
		if xx[i]+yy[j] > bound {
			j = 0
			i++
			if len(xx) == i {
				xx = append(xx, xx[i-1]*x)
			}
		}
	}
	ans := []int{}
	for k := range m {
		ans = append(ans, k)
	}
	return ans
}
