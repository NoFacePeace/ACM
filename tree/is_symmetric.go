package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqual(root.Left, root.Right)
}

func isEqual(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil {
		return false
	}
	if right == nil {
		return false
	}
	return left.Val == right.Val && isEqual(left.Left, right.Right) && isEqual(left.Right, right.Left)
}
