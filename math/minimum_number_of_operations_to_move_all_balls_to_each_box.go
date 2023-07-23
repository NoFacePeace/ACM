package math

func minOperations(boxes string) []int {
	l := len(boxes)
	l1 := make([]int, l)
	l2 := make([]int, l)
	r1 := make([]int, l)
	r2 := make([]int, l)
	for i := 1; i < l; i++ {
		if boxes[i-1] == '1' {
			l1[i] = l1[i-1] + 1
		} else {
			l1[i] = l1[i-1]
		}
		l2[i] = l2[i-1] + l1[i]
	}
	for i := l - 2; i >= 0; i-- {
		if boxes[i+1] == '1' {
			r1[i] = r1[i+1] + 1
		} else {
			r1[i] = r1[i+1]
		}
		r2[i] = r2[i+1] + r1[i]
	}
	ans := make([]int, l)
	for i := 0; i < l; i++ {
		if i == 0 {
			ans[i] = r2[i]
			continue
		}
		if i == l-1 {
			ans[i] = l2[i]
			continue
		}
		ans[i] = l2[i] + r2[i]
	}
	return ans
}
