package array

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}
	carry := 0
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] >= 10 {
			digits[i] -= 10
			carry = 1
			continue
		}
		carry = 0
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
