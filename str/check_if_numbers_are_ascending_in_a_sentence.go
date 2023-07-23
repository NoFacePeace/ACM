package str

import (
	"math"
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	arr := strings.Split(s, " ")
	prev := math.MinInt
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			continue
		}
		if num <= prev {
			return false
		}
		prev = num
	}
	return true
}
