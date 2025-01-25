package main

import "math"

func diameterOfBinaryTree(root *TreeNode) int {
	maxPath := 0

	var calcHeight func(root *TreeNode) int
	calcHeight = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lh := calcHeight(root.Left)
		rh := calcHeight(root.Right)
		path := lh + rh
		maxPath = int(math.Max(float64(maxPath), float64(path)))

		return 1 + int(math.Max(float64(lh), float64(rh)))
	}

	calcHeight(root)
	return maxPath
}
