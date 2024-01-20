package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func robiii(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		l1, l2 := dfs(root.Left)
		r1, r2 := dfs(root.Right)
		m1 := max(l1+r1, l2+r2+root.Val)
		m2 := l1 + r1
		return m1, m2
	}
	return max(dfs(root))
}
