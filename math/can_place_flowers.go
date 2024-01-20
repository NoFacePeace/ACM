package math

func canPlaceFlowers(flowerbed []int, n int) bool {

	cnt := 0
	i := 0
	start := true
	for _, v := range flowerbed {
		if v == 1 {
			if start {
				start = false
				cnt += i / 2
				i = 0
				continue
			}
			cnt += (i - 1) / 2
			i = 0
			continue
		}
		i++
	}
	cnt += i / 2
	if start {
		cnt = (i + 1) / 2
	}
	return cnt >= n
}
