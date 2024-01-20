package tree

func checkTree(root *TreeNode) bool {
	mid := root.Val
	l := root.Left.Val
	r := root.Right.Val
	return mid == l+r
}
