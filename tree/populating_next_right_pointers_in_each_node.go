package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	arr := []*Node{}
	var dfs func(root *Node, index int)
	dfs = func(root *Node, index int) {
		if root == nil {
			return
		}
		if len(arr) <= index {
			arr = append(arr, root)
		}
		tmp := arr[index]
		if tmp != root {
			tmp.Next = root
			arr[index] = root
		}
		dfs(root.Left, index+1)
		dfs(root.Right, index+1)
	}
	dfs(root, 0)
	return root
}
