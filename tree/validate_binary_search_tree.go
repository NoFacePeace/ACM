package tree

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST(root.Left) {
		return false
	}
	if !isValidBST(root.Right) {
		return false
	}
	if root.Left != nil && max(root.Left) >= root.Val {
		return false
	}
	if root.Right != nil && min(root.Right) <= root.Val {
		return false
	}
	return true
}

func min(root *TreeNode) int {
	if root.Left != nil {
		return min(root.Left)
	}
	return root.Val
}
func max(root *TreeNode) int {
	if root.Right != nil {
		return max(root.Right)
	}
	return root.Val
}
