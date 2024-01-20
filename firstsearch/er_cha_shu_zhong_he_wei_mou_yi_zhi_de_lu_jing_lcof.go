package firstsearch

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var dfs func(root *TreeNode, sum int)
	ans := [][]int{}
	arr := []int{}
	dfs = func(root *TreeNode, sum int) {
		sum += root.Val
		if sum == target && root.Left == nil && root.Right == nil {
			arr = append(arr, root.Val)
			tmp := append([]int{}, arr...)
			arr = arr[:len(arr)-1]
			ans = append(ans, tmp)
			return
		}
		arr = append(arr, root.Val)
		if root.Left != nil {
			dfs(root.Left, sum)
		}
		if root.Right != nil {
			dfs(root.Right, sum)
		}
		arr = arr[:len(arr)-1]
	}
	dfs(root, 0)
	return ans
}
