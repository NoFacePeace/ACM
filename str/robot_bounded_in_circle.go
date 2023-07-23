package str

func isRobotBounded(instructions string) bool {
	dir := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	x, y := 0, 0
	cnt := 0
	for i := 0; i < len(instructions); i++ {
		c := instructions[i]
		if c == 'L' {
			cnt += 1
			continue
		}
		if c == 'R' {
			cnt -= 1
			continue
		}
		d := dir[abs(cnt)%4]
		x += d[0]
		y += d[1]
	}
	if x == 0 && y == 0 {
		return true
	}
	return cnt%4 != 0
}
