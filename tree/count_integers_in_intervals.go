package tree

type CountIntervals struct {
	root *item
}

type item struct {
	lv  int
	rv  int
	ll  *item
	rl  *item
	cnt int
}

func Constructor() CountIntervals {
	return CountIntervals{}
}

func (this *CountIntervals) search(root *item, left, right int) int {
	if root.lv <= left && root.rv >= right {
		return 0
	}
	if right <= root.rv {
		cnt := 0
		if right >= root.lv {
			right = root.lv - 1
		}
		if root.ll == nil {
			root.ll = &item{
				lv:  left,
				rv:  right,
				cnt: right - left + 1,
			}
			cnt = right - left + 1
		} else {
			cnt = this.search(root.ll, left, right)
		}
		root.cnt += cnt
		return cnt
	}
	if left >= root.lv {
		cnt := 0
		if left <= root.rv {
			left = root.rv + 1
		}
		if root.rl == nil {
			root.rl = &item{
				lv:  left,
				rv:  right,
				cnt: right - left + 1,
			}
			cnt = right - left + 1
		} else {
			cnt = this.search(root.rl, left, right)
		}
		root.cnt += cnt
		return cnt
	}
	cnt := 0
	if root.ll == nil {
		root.ll = &item{
			lv:  left,
			rv:  root.lv - 1,
			cnt: root.lv - left,
		}
		cnt += root.lv - left
	} else {
		cnt += this.search(root.ll, left, root.lv-1)
	}
	if root.rl == nil {
		root.rl = &item{
			lv:  root.rv + 1,
			rv:  right,
			cnt: right - root.rv,
		}
		cnt += right - root.rv
	} else {
		cnt += this.search(root.rl, root.rv+1, right)
	}
	root.cnt += cnt
	return cnt
}

func (this *CountIntervals) Add(left int, right int) {
	if this.root == nil {
		this.root = &item{
			lv:  left,
			rv:  right,
			cnt: right - left + 1,
		}
		return
	}
	this.search(this.root, left, right)
}

func (this *CountIntervals) Count() int {
	if this.root == nil {
		return 0
	}
	cnt := this.root.cnt
	return cnt
}
