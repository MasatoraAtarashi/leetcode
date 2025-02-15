package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if new, ok := visited[node]; ok {
			return new
		}
		new := &Node{
			Val:       node.Val,
			Neighbors: nil,
		}
		visited[node] = new
		for _, neighbor := range node.Neighbors {
			new.Neighbors = append(new.Neighbors, dfs(neighbor))
		}

		return new
	}

	return dfs(node)
}
