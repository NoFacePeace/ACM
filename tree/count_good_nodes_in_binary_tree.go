package tree

func goodNodes(root *TreeNode) int {
	cnt := 0
	var dfs func(root *TreeNode, max int)
	dfs = func(root *TreeNode, max int) {
		if root == nil {
			return
		}
		if root.Val >= max {
			cnt++
			max = root.Val
		}
		dfs(root.Left, max)
		dfs(root.Right, max)
	}
	dfs(root, root.Val)
	return cnt
}
