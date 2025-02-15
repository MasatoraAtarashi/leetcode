package main

import "fmt"

func buildTree(preorder []int, inorder []int) *TreeNode {
	fmt.Printf("preorder: %d, inorder: %d\n", preorder, inorder)
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := preorder[0]
	var inRootIdx int
	for idx, val := range inorder {
		if val == root {
			inRootIdx = idx
			break
		}
	}

	fmt.Println("root", root)
	fmt.Println("inRootIdx", inRootIdx)

	result := &TreeNode{
		Val:   root,
		Left:  nil,
		Right: nil,
	}

	if len(preorder) > 1 || len(inorder) > 1 {
		result.Left = buildTree(preorder[1:inRootIdx+1], inorder[:inRootIdx+1])
		result.Right = buildTree(preorder[inRootIdx+1:], inorder[inRootIdx+1:])
	}

	return result
}
