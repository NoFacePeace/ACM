package math

func minimumPerimeter(neededApples int64) int64 {
	var sum, edges, cnt int64
	sum = 0
	edges = 0
	cnt = 0
	for sum < neededApples {
		cnt++
		edges = (edges + (cnt*2 + 1 - 2) + cnt*2*2)
		sum += edges*4 - cnt*2*4
	}
	return cnt * 2 * 4
}
