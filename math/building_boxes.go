package math

func minimumBoxes(n int) int {
	sum := 0
	row := 0
	i := 1
	for sum+row+i <= n {
		sum += row + i
		row += i
		i++
	}

	for j := 1; sum+j <= n; j++ {
		sum += j
		row++
	}
	if sum == n {
		return row
	}
	return row + 1
}
