package main

import "fmt"

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0 // leetcode上ではグローバル変数がテストケース間で共通ぽい
	maxPath(root)
	return max
}

func maxPath(root *TreeNode) {
	if root == nil {
		return
	}

	lh := Height(root.Left)
	rh := Height(root.Right)
	fmt.Println(lh, rh)
	if root.Left != nil {
		lh++
	}
	if root.Right != nil {
		rh++
	}

	path := lh + rh
	fmt.Println("path", path)
	if max < (int(path)) {
		max = int(path)
	}

	maxPath(root.Left)
	maxPath(root.Right)

	fmt.Println("max", max)
}
