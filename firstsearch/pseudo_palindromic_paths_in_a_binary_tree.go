package firstsearch

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {
	m := map[int]int{}
	odd, cnt, ans := 0, 0, 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		m[root.Val]++
		cnt++
		if m[root.Val]%2 != 0 {
			odd++
		} else {
			odd--
		}
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
		if root.Left == nil && root.Right == nil {
			if odd == 1 && cnt%2 != 0 {
				ans++
			}
			if odd == 0 && cnt%2 == 0 {
				ans++
			}
		}
		m[root.Val]--
		if m[root.Val]%2 != 0 {
			odd++
		} else {
			odd--
		}
		cnt--
	}
	if root == nil {
		return 0
	}
	dfs(root)
	return ans
}
