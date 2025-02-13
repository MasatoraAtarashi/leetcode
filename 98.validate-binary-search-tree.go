package main

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	result := true
	var traverse98 func(t *TreeNode, min, max int)
	traverse98 = func(t *TreeNode, min, max int) {
		if t.Val <= min || max <= t.Val {
			result = false
			return
		}

		if t.Left != nil {
			traverse98(t.Left, min, t.Val)
		}
		if t.Right != nil {
			traverse98(t.Right, t.Val, max)
		}
	}

	traverse98(root, -999999999999999999, 999999999999999999)
	return result
}
