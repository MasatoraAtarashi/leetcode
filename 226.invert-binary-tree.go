package main

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root = Swap(root)
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)

	return root
}
