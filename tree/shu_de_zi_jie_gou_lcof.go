package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var equal func(a, b *TreeNode) bool
	equal = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil {
			return false
		}
		if b == nil {
			return true
		}
		if a.Val != b.Val {
			return false
		}
		return equal(a.Left, b.Left) && equal(a.Right, b.Right)
	}
	var f func(a, b *TreeNode) bool
	f = func(a, b *TreeNode) bool {
		if a == nil {
			return false
		}
		if b == nil {
			return false
		}
		if equal(a, b) {
			return true
		}
		return f(a.Left, b) || f(a.Right, b)
	}
	return f(A, B)
}
