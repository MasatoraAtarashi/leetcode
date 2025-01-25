package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Swap(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}

	return &TreeNode{
		Val:   tree.Val,
		Left:  tree.Right,
		Right: tree.Left,
	}
}
