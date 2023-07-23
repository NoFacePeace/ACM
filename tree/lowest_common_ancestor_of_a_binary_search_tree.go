package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	var ans *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root == p || root == q {
			return root
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		if l == p && r == q {
			ans = root
			return root
		}
		if l == q && r == p {
			ans = root
			return root
		}
		if l == q || l == p {
			return l
		}
		if r == q || r == p {
			return r
		}
		return root
	}
	val := dfs(root)
	if val == p || val == q {
		return val
	}
	return ans
}
