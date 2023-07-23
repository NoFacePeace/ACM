package array

func countPoints(points [][]int, queries [][]int) []int {
	ans := []int{}
	square := func(a int) int {
		return a * a
	}
	for _, q := range queries {
		cnt := 0
		for _, p := range points {
			if square(q[0]-p[0])+square(q[1]-p[1]) <= square(q[2]) {
				cnt++
			}
		}
		ans = append(ans, cnt)
	}
	return ans
}
