package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for queue.Length() > 0 {
		// 同じレベルのNodeを処理
		levelSize := queue.Length()
		for i := 0; i < levelSize; i++ {
			cur := queue.Dequeue()
			if cur.Val.Left != nil {
				queue.Enqueue(cur.Val.Left)
			}
			if cur.Val.Right != nil {
				queue.Enqueue(cur.Val.Right)
			}
			if i == levelSize-1 {
				result = append(result, cur.Val.Val)
			}
		}
	}

	return result
}
