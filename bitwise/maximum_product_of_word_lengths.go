package bitwise

func maxProduct(words []string) int {
	n := len(words)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		word := words[i]
		num := 0
		for _, c := range word {
			num = num | 1<<int(c-'a')
		}
		nums[i] = num
	}
	max := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]&nums[j] != 0 {
				continue
			}
			l := len(words[i]) * len(words[j])
			if l > max {
				max = l
			}
		}
	}
	return max
}
