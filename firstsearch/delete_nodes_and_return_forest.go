package firstsearch

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	m := map[int]bool{}
	for _, v := range to_delete {
		m[v] = true
	}
	ans := []*TreeNode{}
	q := []*TreeNode{root}
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		val := root.Val
		if _, ok := m[val]; ok {
			if root.Left != nil {
				q = append(q, root.Left)
			}
			if root.Right != nil {
				q = append(q, root.Right)
			}
			return true
		}
		if dfs(root.Left) {
			root.Left = nil
		}
		if dfs(root.Right) {
			root.Right = nil
		}
		return false
	}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if !dfs(node) {
			ans = append(ans, node)
		}
	}
	return ans
}
