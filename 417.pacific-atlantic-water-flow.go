package main

func pacificAtlantic(heights [][]int) [][]int {
	var result [][]int
	added := make([][]bool, len(heights))
	pacific := make([][]bool, len(heights))
	var fs func(ocean string, row, col, before int)
	fs = func(ocean string, row, col, before int) {
		// fmt.Println(row, col, before)
		if row >= len(heights) || row < 0 || col >= len(heights[row]) || col < 0 {
			return
		}
		now := heights[row][col]
		if now < before {
			return
		}

		if ocean == "p" {
			if pacific[row] == nil {
				pacific[row] = make([]bool, len(heights[row]))
			}
			pacific[row][col] = true

			fs(ocean, row+1, col, now)
			fs(ocean, row, col+1, now)
		} else {
			if pacific[row] != nil {
				if pacific[row][col] {
					if added[row] == nil || !added[row][col] {
						result = append(result, []int{row, col})
					}
					if added[row] == nil {
						added[row] = make([]bool, len(heights[row]))
					}
					added[row][col] = true
				}
			}
			fs(ocean, row-1, col, now)
			fs(ocean, row, col-1, now)
		}
	}

	for row := range len(heights) {
		fs("p", row, 0, 0)
	}
	for col := range len(heights[0]) {
		fs("p", 0, col, 0)
	}
	for row := range len(heights) {
		fs("a", row, len(heights)-1, 0)
	}
	for col := range len(heights[len(heights)-1]) {
		fs("a", len(heights[len(heights)-1])-1, col, 0)
	}

	// for i := range pacific {
	// 	for j := range pacific[i] {
	// 		fmt.Println(heights[i][j], pacific[i][j])
	// 	}
	// }

	return result
}
