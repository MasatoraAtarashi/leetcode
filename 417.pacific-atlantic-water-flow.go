package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
	var result [][]int
	added := make([][]bool, len(heights))
	pacific := make([][]bool, len(heights))
	atlantic := make([][]bool, len(heights))
	var dfs func(ocean string, row, col, before int)
	dfs = func(ocean string, row, col, before int) {
		if row >= len(heights) || row < 0 || col >= len(heights[row]) || col < 0 {
			return
		}
		fmt.Printf("Visiting %s: (%d,%d) height=%d, prev=%d\n",
			ocean, row, col, heights[row][col], before)

		now := heights[row][col]
		if now < before {
			return
		}

		if ocean == "p" {
			if pacific[row] == nil {
				pacific[row] = make([]bool, len(heights[row]))
			}
			if pacific[row][col] {
				return
			}
			pacific[row][col] = true
		} else {
			if atlantic[row] == nil {
				atlantic[row] = make([]bool, len(heights[row]))
			}
			if atlantic[row][col] {
				return
			}
			atlantic[row][col] = true
			if pacific[row] != nil {
				if pacific[row][col] {
					if added[row] == nil || !added[row][col] {
						fmt.Printf("Adding to result: (%d,%d)\n", row, col)
						result = append(result, []int{row, col})
					}
					if added[row] == nil {
						added[row] = make([]bool, len(heights[row]))
					}
					added[row][col] = true
				}
			}
		}
		dfs(ocean, row+1, col, now)
		dfs(ocean, row, col+1, now)
		dfs(ocean, row-1, col, now)
		dfs(ocean, row, col-1, now)
	}

	fmt.Println("\n=== Pacific Left Edge ===")
	for row := range len(heights) {
		fmt.Printf("Starting from left: row=%d\n", row)
		dfs("p", row, 0, 0)
	}

	fmt.Println("\n=== Pacific Top Edge ===")
	for col := range len(heights[0]) {
		fmt.Printf("Starting from top: col=%d\n", col)
		dfs("p", 0, col, 0)
	}

	fmt.Println("\n=== Atlantic Right Edge ===")
	for row := range len(heights) {
		fmt.Printf("Starting from right: row=%d\n", row)
		dfs("a", row, len(heights[row])-1, 0)
	}

	fmt.Println("\n=== Atlantic Bottom Edge ===")
	for col := range len(heights[len(heights)-1]) {
		fmt.Printf("Starting from bottom: col=%d, row=%d\n",
			col, len(heights)-1)
		dfs("a", len(heights)-1, col, 0)
	}

	return result
}
