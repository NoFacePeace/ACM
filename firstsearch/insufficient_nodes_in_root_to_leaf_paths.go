package firstsearch

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(root *TreeNode, sum int) bool
	dfs = func(root *TreeNode, sum int) bool {
		if root.Left == nil && root.Right == nil {
			return root.Val+sum >= limit
		}
		lok := false
		rok := false
		if root.Left != nil && dfs(root.Left, sum+root.Val) {
			lok = true
		} else {
			root.Left = nil
		}
		if root.Right != nil && dfs(root.Right, sum+root.Val) {
			rok = true
		} else {
			root.Right = nil
		}
		if lok || rok {
			return true
		}
		return false
	}
	if root == nil {
		return root
	}
	if dfs(root, 0) {
		return root
	}
	return nil
}
