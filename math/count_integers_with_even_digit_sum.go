package math

func countEven(num int) int {
	ans := 0
	for i := 1; i <= num; i++ {
		j := i
		odd := 0
		even := 0
		for j != 0 {
			mod := j % 10
			j /= 10
			if mod%2 == 0 {
				even++
			} else {
				odd++
			}
		}
		if (odd+even)%2 == 0 {
			ans++
		}
	}
	return ans
}
