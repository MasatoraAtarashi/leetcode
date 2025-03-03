package main

import "fmt"

const (
	FRESH = 1
	ROTT  = 2
)

type coordinate struct {
	x int
	y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

func orangesRotting(grid [][]int) int {
	// 最初に新鮮なオレンジの数を数え、腐ったオレンジをキューをにいれる
	rottQueue := Queue[coordinate]{}
	var freshCount int
	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == FRESH {
				freshCount++
				continue
			}
			if grid[i][j] == ROTT {
				rottQueue.Enqueue(coordinate{i, j})
			}
		}
	}

	var time int
	var rottaize func(i, j int)
	rottaize = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
			return
		}
		if grid[i][j] == FRESH {
			freshCount--
			rottQueue.Enqueue(coordinate{i, j})
			grid[i][j] = ROTT
		}
	}

	for rottQueue.Length() != 0 && freshCount != 0 {
		length := rottQueue.Length()
		for range length {
			r := rottQueue.Dequeue()

			// 上
			rottaize(r.Val.x+1, r.Val.y)
			// 左
			rottaize(r.Val.x, r.Val.y-1)
			// 右
			rottaize(r.Val.x, r.Val.y+1)
			// 下
			rottaize(r.Val.x-1, r.Val.y)
		}
		time++
	}

	if freshCount == 0 {
		return time
	}
	return -1
}
