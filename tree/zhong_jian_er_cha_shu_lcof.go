package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTreeLcof(preorder []int, inorder []int) *TreeNode {
	var f func(preorder []int, inorder []int) *TreeNode
	f = func(preorder []int, inorder []int) *TreeNode {
		n := len(preorder)
		if n == 0 {
			return nil
		}
		val := preorder[0]
		i := 0
		for i < n {
			if inorder[i] == val {
				break
			}
			i++
		}
		root := &TreeNode{
			Val: val,
		}
		root.Left = f(preorder[1:1+i], inorder[:i])
		root.Right = f(preorder[1+i:], inorder[i+1:])
		return root
	}
	return f(preorder, inorder)
}
