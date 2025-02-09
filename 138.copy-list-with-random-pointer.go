package main

func copyRandomList(head *Node) *Node {
	// step1 ノードをコピーしてマッピング
	nodeMap := make(map[*Node]*Node, 0)
	old := head
	for old != nil {
		nodeMap[old] = &Node{
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

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
