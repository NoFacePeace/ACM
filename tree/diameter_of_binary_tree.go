package tree

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	ans := 0
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		if l+r > ans {
			ans = l + r
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	dfs(root)
	return ans
}
