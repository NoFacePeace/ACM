package hash

func countBalls(lowLimit int, highLimit int) int {
	m := map[int]int{}
	convert := func(n int) int {
		sum := 0
		for {
			sum += n % 10
			n /= 10
			if n == 0 {
				break
			}
		}
		return sum
	}
	ans := 0
	for i := lowLimit; i <= highLimit; i++ {
		num := convert(i)
		m[num]++
		if m[num] > ans {
			ans = m[num]
		}
	}
	return ans
}
