package math

import "strconv"

func fizzBuzz(n int) []string {
	name := "FizzBuzz"
	first := "Fizz"
	last := "Buzz"
	ans := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, name)
			continue
		}
		if i%3 == 0 {
			ans = append(ans, first)
			continue
		}
		if i%5 == 0 {
			ans = append(ans, last)
			continue
		}
		ans = append(ans, strconv.Itoa(i))
	}
	return ans
}
