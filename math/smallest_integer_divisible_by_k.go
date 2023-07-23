package math

func smallestRepunitDivByK(k int) int {
	m := map[int]bool{}
	num := 1
	cnt := 1
	for {
		for num < k {
			num = num*10 + 1
			cnt++
		}
		num = num % k
		if num == 0 {
			return cnt
		}
		if m[num] {
			break
		}
		m[num] = true
	}
	return -1
}
