package math

func findSolution(customFunction func(int, int) int, z int) [][]int {
	x := 1
	y := 1000
	ans := [][]int{}
	for x <= 1000 && y >= 1 {
		val := customFunction(x, y)
		if val == z {
			ans = append(ans, []int{x, y})
			x++
			y--
			continue
		}
		if val < z {
			x++
			continue
		}
		y--
	}
	return ans
}
