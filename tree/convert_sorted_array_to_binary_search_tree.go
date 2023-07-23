package tree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	l := 0
	r := len(nums) - 1
	mid := (l + r) / 2
	node := &TreeNode{
		Val: nums[mid],
	}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}
