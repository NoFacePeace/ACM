package tree

type NumArray struct {
	arr []int
	n   int
}

func NewNumArray(nums []int) NumArray {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	arr := make([]int, 4*n)
	var bk func(i, l, r int)
	bk = func(i, l, r int) {
		if l == 0 {
			arr[i] = nums[r]
		} else {
			arr[i] = nums[r] - nums[l-1]
		}
		if l == r {
			return
		}
		mid := (l + r) / 2
		bk(2*i+1, l, mid)
		bk(2*i+2, mid+1, r)
	}
	bk(0, 0, n-1)
	return NumArray{
		arr: arr,
		n:   n - 1,
	}
}

func (this *NumArray) Update(index int, val int) {
	this.update(0, 0, this.n, index, val)
}

func (this *NumArray) update(i, l, r, index, val int) {
	if l == r && l == index {
		this.arr[i] = val
		return
	}
	mid := (l + r) / 2
	if index <= mid {
		this.update(2*i+1, l, mid, index, val)
	} else {
		this.update(2*i+2, mid+1, r, index, val)
	}
	this.arr[i] = this.arr[2*i+1] + this.arr[2*i+2]
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(0, 0, this.n, left, right)
}

func (this *NumArray) query(idx, i, j, left, right int) int {
	if left <= i && right >= j {
		return this.arr[idx]
	}
	ans := 0
	mid := (i + j) / 2
	if left <= mid {
		ans += this.query(idx*2+1, i, mid, left, right)
	}
	if right > mid {
		ans += this.query(idx*2+2, mid+1, j, left, right)
	}
	return ans
}
