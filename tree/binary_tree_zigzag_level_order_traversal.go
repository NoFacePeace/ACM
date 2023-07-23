package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	var dfs func(root *TreeNode, index int)
	dfs = func(root *TreeNode, index int) {
		if root == nil {
			return
		}
		if len(ans) <= index {
			ans = append(ans, []int{root.Val})
		} else {
			if index%2 == 0 {
				ans[index] = append(ans[index], root.Val)
			} else {
				ans[index] = append([]int{root.Val}, ans[index]...)
			}
		}
		dfs(root.Left, index+1)
		dfs(root.Right, index+1)
	}
	dfs(root, 0)
	return ans
}
