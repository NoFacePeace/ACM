package firstsearch

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		values := []int{}
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			values = append(values, v.Val)
		}
		if level%2 == 0 {
			level++
			continue
		}
		i := len(values) - 1
		for _, v := range tmp {
			v.Val = values[i]
			i--
		}
		level++
	}
	return root
}
