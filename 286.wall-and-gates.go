package main

const (
	WALL = -1
	GATE = 0
	INF  = 2147483647
)

func wallsAndGates(rooms [][]int) {
	gateQueue := Queue[coordinate]{}
	var infCount int

	// ゲートをキューにいれる & 埋まってないセルの数を数える
	for i := range len(rooms) {
		for j := range len(rooms[i]) {
			if rooms[i][j] == GATE {
				gateQueue.Enqueue(coordinate{i, j})
			}
			if rooms[i][j] == INF {
				infCount++
			}
		}
	}

	distance := 1
	next := func(i, j int) {
		if i < 0 || i >= len(rooms) || j < 0 || j >= len(rooms[i]) {
			return
		}
		if rooms[i][j] == INF {
			rooms[i][j] = distance
			gateQueue.Enqueue(coordinate{i, j})
		}
	}

	for gateQueue.Length() != 0 && infCount != 0 {
		length := gateQueue.Length()
		for range length {
			g := gateQueue.Dequeue()
			// 上
			next(g.Val.x-1, g.Val.y)
			// 左
			next(g.Val.x, g.Val.y-1)
			// 右
			next(g.Val.x, g.Val.y+1)
			// 下
			next(g.Val.x+1, g.Val.y)
		}
		distance++
	}
}
