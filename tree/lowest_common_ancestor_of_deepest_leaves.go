package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var f func(root *TreeNode, depth int) int
	var maxDepth, ansDepth int
	max, ans := root, root
	f = func(root *TreeNode, depth int) int {
		if root.Left == nil && root.Right == nil {
			if depth > maxDepth {
				maxDepth = depth
				max = root
			}
			return depth
		}
		if root.Left == nil {
			return f(root.Right, depth+1)
		}
		if root.Right == nil {
			return f(root.Left, depth+1)
		}
		l := f(root.Left, depth+1)
		r := f(root.Right, depth+1)
		if l == r && l >= maxDepth {
			ans = root
			ansDepth = l
		}
		if l > r {
			return l
		}
		return r
	}
	f(root, 0)
	if maxDepth == ansDepth {
		return ans
	}
	return max
}
