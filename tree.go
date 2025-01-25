package main

import "math"

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

func Height(tree *TreeNode) float64 {
	if tree == nil {
		return 0
	}

	if tree.Left == nil && tree.Right == nil {
		return 0
	}

	lh := Height(tree.Left)
	rh := Height(tree.Right)

	if tree.Left != nil {
		lh++
	}
	if tree.Right != nil {
		rh++
	}

	return math.Max(lh, rh)
}
