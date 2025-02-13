package main

func kthSmallest(root *TreeNode, k int) int {
	var count int
	var result int
	var traverseInOrder func(t *TreeNode)
	traverseInOrder = func(t *TreeNode) {
		if t.Left != nil {
			traverseInOrder(t.Left)
		}
		count++
		if count == k {
			result = t.Val
			return
		}
		if t.Right != nil {
			traverseInOrder(t.Right)
		}
	}
	traverseInOrder(root)
	return result
}
