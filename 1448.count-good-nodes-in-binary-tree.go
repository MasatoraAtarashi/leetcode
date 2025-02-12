package main

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var pathMaxCnt int

	stack := Stack[*TreeNode]{}
	var traverse func(t *TreeNode, pathMax int)
	traverse = func(t *TreeNode, pathMax int) {
		stack.Push(t)
		if t.Val >= pathMax {
			pathMax = t.Val
			pathMaxCnt++
		}
		if t.Left != nil {
			traverse(t.Left, pathMax)
		}
		if t.Right != nil {
			traverse(t.Right, pathMax)
		}
		stack.Pop()
	}

	traverse(root, root.Val)
	return pathMaxCnt
}
