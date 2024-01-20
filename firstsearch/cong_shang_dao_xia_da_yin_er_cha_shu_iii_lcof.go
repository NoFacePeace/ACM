package firstsearch

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	q := []*TreeNode{root}
	d := true
	for len(q) > 0 {
		tmp := q
		q = nil
		arr := []int{}
		for _, v := range tmp {
			arr = append(arr, v.Val)
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}
		if !d {
			n := len(arr)
			for i := 0; i < n/2; i++ {
				arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
			}
		}
		d = !d
		ans = append(ans, arr)
	}
	return ans
}
