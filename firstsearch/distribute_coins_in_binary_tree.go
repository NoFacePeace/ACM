package firstsearch

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	move := 0
	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		move += abs(l)
		r := dfs(root.Right)
		move += abs(r)
		need := l + r
		if root.Val == 0 {
			need -= 1
		} else {
			need += root.Val - 1
		}
		return need
	}
	dfs(root)
	return move
}
