package main

type ocean struct {
	visited   [][]bool
	reachable [][]bool
}

func pacificAtlantic(heights [][]int) [][]int {
	var result [][]int
	pacific := ocean{
		visited:   make([][]bool, len(heights)),
		reachable: make([][]bool, len(heights)),
	}
	atlantic := ocean{
		visited:   make([][]bool, len(heights)),
		reachable: make([][]bool, len(heights)),
	}
	var fs func(ocean *ocean, row, col, before int)
	fs = func(ocean *ocean, row, col, before int) {
		// fmt.Println(row, col, before)
		if row >= len(heights) || row < 0 || col >= len(heights[row]) || col < 0 {
			return
		}
		now := heights[row][col]
		if now < before {
			return
		}

		if ocean.visited[row] != nil {
			if ocean.visited[row][col] {
				return
			}
		}

		if ocean.visited[row] == nil {
			ocean.visited[row] = make([]bool, len(heights[row]))
		}
		ocean.visited[row][col] = true

		if ocean.reachable[row] == nil {
			ocean.reachable[row] = make([]bool, len(heights[row]))
		}
		ocean.reachable[row][col] = true
		if pacific.reachable[row] != nil && atlantic.reachable[row] != nil {
			if pacific.reachable[row][col] && atlantic.reachable[row][col] {
				result = append(result, []int{row, col})
			}
		}
		fs(ocean, row+1, col, now)
		fs(ocean, row, col+1, now)
		fs(ocean, row-1, col, now)
		fs(ocean, row, col-1, now)
	}

	for row := range len(heights) {
		pacific.visited = make([][]bool, len(heights))
		fs(&pacific, row, 0, 0)
	}
	for col := range len(heights[0]) {
		pacific.visited = make([][]bool, len(heights))
		fs(&pacific, 0, col, 0)
	}
	for row := range len(heights) {
		atlantic.visited = make([][]bool, len(heights))
		fs(&atlantic, row, len(heights)-1, 0)
	}
	for col := range len(heights[len(heights)-1]) {
		atlantic.visited = make([][]bool, len(heights))
		fs(&atlantic, len(heights[len(heights)-1]), col, 0)
	}

	// for i := range pacific {
	// 	for j := range pacific[i] {
	// 		fmt.Println(heights[i][j], pacific[i][j])
	// 	}
	// }

	return result
}
