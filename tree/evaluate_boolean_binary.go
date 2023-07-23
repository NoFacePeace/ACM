package tree

func evaluateTree(root *TreeNode) bool {
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Left == nil || root.Right == nil {
			return root.Val == 1
		}
		if root.Val == 2 {
			return dfs(root.Left) || dfs(root.Right)
		}
		return dfs(root.Left) && dfs(root.Right)
	}
	return dfs(root)
}
