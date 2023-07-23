package firstsearch

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	arr := []int{}
	ans := 0
	abs := func(a, b int) int {
		if a < b {
			return b - a
		}
		return a - b
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		arr = append(arr, root.Val)
		i := len(arr)
		dfs(root.Left)
		dfs(root.Right)
		min := root.Val
		max := root.Val
		for i < len(arr) {
			if arr[i] > max {
				max = arr[i]
			}
			if arr[i] < min {
				min = arr[i]
			}
			i++
		}
		if abs(root.Val, max) > ans {
			ans = abs(root.Val, max)
		}
		if abs(root.Val, min) > ans {
			ans = abs(root.Val, min)
		}
	}
	dfs(root)
	return ans
}
