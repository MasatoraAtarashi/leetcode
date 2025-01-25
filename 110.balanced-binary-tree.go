package main

import (
	"math"
)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := Height(root.Left)
	rh := Height(root.Right)
	if root.Left != nil {
		lh++
	}
	if root.Right != nil {
		rh++
	}

	if math.Abs(lh-rh) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
