package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{
		Val: preorder[0],
	}
	if len(preorder) == 1 {
		return node
	}
	pivot := preorder[0]
	i := 0
	for i < len(inorder) && inorder[i] != pivot {
		i++
	}
	node.Left = buildTree(preorder[1:i+1], inorder[:i])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return node
}
