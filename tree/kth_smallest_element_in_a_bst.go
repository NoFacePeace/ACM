package tree

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(root *TreeNode)
	arr := []int{}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		arr = append(arr, root.Val)
		if len(arr) == k {
			return
		}
		dfs(root.Right)
	}
	dfs(root)
	if len(arr) >= k {
		return arr[k-1]
	}
	return -1
}
