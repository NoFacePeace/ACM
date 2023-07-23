package tree

type Node1 struct {
	Val      int
	Children []*Node1
}

func preorder(root *Node1) []int {
	ans := []int{}
	var preorder func(root *Node1)
	preorder = func(root *Node1) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		for _, v := range root.Children {
			preorder(v)
		}
	}
	preorder(root)
	return ans
}
