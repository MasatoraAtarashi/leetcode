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

func IsSameTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return IsSameTree(t1.Left, t2.Left) && IsSameTree(t1.Right, t2.Right)
}
