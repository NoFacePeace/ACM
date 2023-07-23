package array

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	sum := 0
	two := 0
	for _, v := range colsum {
		sum += v
		if v == 2 {
			two++
		}
	}
	if sum != upper+lower {
		return [][]int{}
	}
	if upper < two {
		return [][]int{}
	}
	if lower < two {
		return [][]int{}
	}
	ans := make([][]int, 2)
	n := len(colsum)
	ans[0] = make([]int, n)
	ans[1] = make([]int, n)
	for i := 0; i < n; i++ {
		if colsum[i] == 0 {
			continue
		}
		if colsum[i] == 2 {
			ans[0][i] = 1
			ans[1][i] = 1
			upper--
			lower--
			two--
			continue
		}
		if upper > 0 && upper > two {
			ans[0][i] = 1
			upper--
			continue
		}
		if lower > 0 && lower > two {
			ans[1][i] = 1
			lower--
			continue
		}
		return [][]int{}
	}
	return ans
}
