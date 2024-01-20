package array

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	ans := [][]int{}
	var i int
	for i = 0; i < n; i++ {
		if intervals[i][1] < newInterval[0] {
			ans = append(ans, intervals[i])
			continue
		}
		if intervals[i][0] > newInterval[1] {
			ans = append(ans, newInterval)
			ans = append(ans, intervals[i:]...)
			break
		}
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
	}
	if i == n {
		ans = append(ans, newInterval)
	}
	return ans
}
