package tree

func pathSum(root *TreeNode, targetSum int) int {
	ans := 0
	var f func(root *TreeNode) []int
	f = func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		l := f(root.Left)
		r := f(root.Right)
		arr := []int{}
		arr = append(arr, l...)
		arr = append(arr, r...)
		for i := 0; i < len(arr); i++ {
			arr[i] += root.Val
			if arr[i] == targetSum {
				ans++
			}
		}
		if root.Val == targetSum {
			ans++
		}
		arr = append(arr, root.Val)
		return arr
	}
	f(root)
	return ans
}
