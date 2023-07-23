package firstsearch

func maxSumBST(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) (bool, int, int, int)
	dfs = func(root *TreeNode) (bool, int, int, int) {
		if root.Left == nil && root.Right == nil {
			if root.Val > ans {
				ans = root.Val
			}
			return true, root.Val, root.Val, root.Val
		}
		if root.Left == nil {
			ok, min, max, sum := dfs(root.Right)
			if !ok || root.Val >= min {
				root.Right = nil
				return false, min, max, sum
			}
			if sum+root.Val > ans {
				ans = sum + root.Val
			}
			return true, root.Val, max, sum + root.Val
		}
		if root.Right == nil {
			ok, min, max, sum := dfs(root.Left)
			if !ok || root.Val <= max {
				root.Left = nil
				return false, min, max, sum
			}
			if sum+root.Val > ans {
				ans = sum + root.Val
			}
			return true, min, root.Val, sum + root.Val
		}
		lok, lmin, lmax, lsum := dfs(root.Left)
		rok, rmin, rmax, rsum := dfs(root.Right)
		if !lok || root.Val <= lmax {
			return false, lmin, lmax, lsum
		}
		if !rok || root.Val >= rmin {
			return false, rmin, rmax, rsum
		}
		if lsum+root.Val+rsum > ans {
			ans = lsum + root.Val + rsum
		}
		return true, lmin, rmax, lsum + root.Val + rsum
	}
	dfs(root)
	return ans
}
