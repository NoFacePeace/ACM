package array

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	size := 8
	m := make([][]int, size)
	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
	}
	for _, queen := range queens {
		x, y := queen[0], queen[1]
		m[x][y] = 1
	}
	kx, ky := king[0], king[1]
	ans := [][]int{}
	var f func(x, y, cnt int)
	f = func(x, y, cnt int) {
		if kx+x*cnt < 0 {
			return
		}
		if kx+x*cnt >= size {
			return
		}
		if ky+y*cnt < 0 {
			return
		}
		if ky+y*cnt >= size {
			return
		}
		if m[kx+x*cnt][ky+y*cnt] == 1 {
			ans = append(ans, []int{kx + x*cnt, ky + y*cnt})
			return
		}
		f(x, y, cnt+1)
	}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			f(i, j, 1)
		}
	}
	return ans
}
