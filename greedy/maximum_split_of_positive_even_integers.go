package greedy

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	cnt := finalSum / 2
	if cnt == 1 {
		return []int64{2}
	}
	ans := []int64{}
	for i := int64(1); i <= cnt; i++ {
		cnt -= i
		ans = append(ans, i*2)
		if cnt <= i {
			ans[len(ans)-1] += cnt * 2
			break
		}
	}
	return ans
}
