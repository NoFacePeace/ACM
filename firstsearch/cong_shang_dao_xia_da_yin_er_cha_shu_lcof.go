package firstsearch

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	ans := []int{}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		ans = append(ans, x.Val)
		if x.Left != nil {
			q = append(q, x.Left)
		}
		if x.Right != nil {
			q = append(q, x.Right)
		}
	}
	return ans
}
