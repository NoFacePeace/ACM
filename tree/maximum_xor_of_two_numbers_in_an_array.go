package tree

func findMaximumXOR(nums []int) int {
	root := &trie{}
	x := 0
	for i := 1; i < len(nums); i++ {
		root.add(nums[i-1])
		x = max1(x, root.check(nums[i]))
	}
	return x
}

type trie struct {
	left, right *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := 30; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}

func (t *trie) check(num int) (x int) {
	cur := t
	for i := 30; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1
			} else {
				cur = cur.left
				x = x * 2
			}
		} else {
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
