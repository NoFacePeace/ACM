package array

import "math/rand"

type Solution struct {
	nums     []int
	original []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums:     nums,
		original: append([]int{}, nums...),
	}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	l := len(this.nums)
	for i := 0; i < l; i++ {
		j := i + rand.Intn(l-i)
		this.nums[i], this.nums[j] = this.nums[j], this.nums[i]
	}
	return this.nums
}
