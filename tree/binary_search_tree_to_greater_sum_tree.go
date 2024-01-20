package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	arr := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	sum := 0
	for _, v := range arr {
		sum += v
	}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		sum -= root.Val
		root.Val += sum
		dfs(root.Right)
	}
	dfs(root)
	return root
}
