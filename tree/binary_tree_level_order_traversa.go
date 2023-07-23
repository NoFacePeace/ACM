package tree

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	var dfs func(root *TreeNode, index int)
	dfs = func(root *TreeNode, index int) {
		if root == nil {
			return
		}
		if len(ans) <= index {
			ans = append(ans, []int{root.Val})
		} else {
			ans[index] = append(ans[index], root.Val)
		}
		dfs(root.Left, index+1)
		dfs(root.Right, index+1)
	}
	dfs(root, 0)
	return ans
}
