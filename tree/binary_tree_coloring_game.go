package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root.Val == x {
			return root
		}
		node := dfs(root.Left)
		if node != nil {
			return node
		}
		node = dfs(root.Right)
		return node
	}
	xNode := dfs(root)
	var cnt func(root *TreeNode) int
	cnt = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return 1 + cnt(root.Left) + cnt(root.Right)
	}
	l := cnt(xNode.Left)
	r := cnt(xNode.Right)
	h := n - l - r - 1
	if h > n-h {
		return true
	}
	if l > h+1+r {
		return true
	}
	if r > h+1+l {
		return true
	}
	return false
}
