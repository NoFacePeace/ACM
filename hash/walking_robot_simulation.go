package hash

func robotSim(commands []int, obstacles [][]int) int {
	type coordinate struct {
		x int
		y int
	}
	m := map[coordinate]bool{}
	for _, v := range obstacles {
		c := coordinate{
			x: v[0],
			y: v[1],
		}
		m[c] = true
	}
	x := 0
	y := 0
	dir := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	d := 1
	for _, v := range commands {
		if v == -2 {
			d++
			d %= 4
			continue
		}
		if v == -1 {
			d = (d + 4 - 1) % 4
			continue
		}
		for i := 0; i < v; i++ {
			c := coordinate{
				x: x + dir[d][0],
				y: y + dir[d][1],
			}
			if m[c] {
				break
			}
			x = c.x
			y = c.y
		}
	}
	return x*x + y*y
}
