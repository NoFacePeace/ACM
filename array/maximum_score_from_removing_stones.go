package array

func maximumScore(a int, b int, c int) int {
	sort := func(num []int) []int {
		num1 := []int{}
		num1 = append(num1, num...)
		for i := len(num1); i > 0; i-- {
			for j := 0; j < i-1; j++ {
				if num1[j] > num1[j+1] {
					num1[j], num1[j+1] = num1[j+1], num1[j]
				}
			}
		}
		return num1
	}
	arr := []int{a, b, c}
	arr = sort(arr)
	ans := 0
	for arr[1] != 0 {
		arr[1]--
		arr[2]--
		arr = sort(arr)
		ans++
	}
	return ans
}
