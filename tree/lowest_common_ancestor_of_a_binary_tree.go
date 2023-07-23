package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(n *TreeNode) *TreeNode
	dfs = func(n *TreeNode) *TreeNode {
		if n == nil {
			return nil
		}
		if n == p || n == q {
			return n
		}
		l := dfs(n.Left)
		r := dfs(n.Right)
		if l != nil && r != nil {
			return n
		}
		if l == nil {
			return r
		}
		return l
	}
	return dfs(root)
}
