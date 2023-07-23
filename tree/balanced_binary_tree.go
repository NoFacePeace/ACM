package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	abs := func(a, b int) int {
		if a < b {
			return b - a
		}
		return a - b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var f func(root *TreeNode) (bool, int)
	f = func(root *TreeNode) (bool, int) {
		if root == nil {
			return true, 0
		}
		ok, l := f(root.Left)
		if !ok {
			return false, 0
		}
		ok, r := f(root.Right)
		if !ok {
			return false, 0
		}
		dist := abs(l, r)
		if dist > 1 {
			return false, 0
		}
		return true, max(l, r) + 1
	}
	ok, _ := f(root)
	return ok
}
