package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val == p.Val || root.Val == q.Val {
			return root
		}
		if p.Val < root.Val && root.Val < q.Val {
			return root
		}
		if q.Val < root.Val && root.Val < p.Val {
			return root
		}

		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return nil
}
