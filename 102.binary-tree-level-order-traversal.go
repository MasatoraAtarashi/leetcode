package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := NewQueue[*TreeNode]()
	queue.Enqueue(root)

	for queue.Length() > 0 {
		// 同じレベルにあるノードを処理
		levelSize := queue.Length()
		var level []int

		for i := 0; i < levelSize; i++ {
			current := queue.Dequeue()
			level = append(level, current.Val.Val)
			if current.Val.Left != nil {
				queue.Enqueue(current.Val.Left)
			}
			if current.Val.Right != nil {
				queue.Enqueue(current.Val.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
