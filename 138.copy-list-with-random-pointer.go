package main

func copyRandomList(head *LinkedListNode) *LinkedListNode {
	// step1 ノードをコピーしてマッピング
	nodeMap := make(map[*LinkedListNode]*LinkedListNode, 0)
	old := head
	for old != nil {
		nodeMap[old] = &LinkedListNode{
			Val: old.Val,
		}
		old = old.Next
	}

	// step2 接続(next, random)
	old = head
	for old != nil {
		nodeMap[old].Next = nodeMap[old.Next]
		nodeMap[old].Random = nodeMap[old.Random]
		old = old.Next
	}

	return nodeMap[head]
}

type LinkedListNode struct {
	Val    int
	Next   *LinkedListNode
	Random *LinkedListNode
}
