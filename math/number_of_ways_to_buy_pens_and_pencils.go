package math

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	ans := 0
	m := total / cost1
	for i := 0; i <= m; i++ {
		ans += (total-i*cost1)/cost2 + 1
	}
	return int64(ans)
}
