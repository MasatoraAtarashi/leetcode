package main

func maxAreaOfIsland(grid [][]int) int {
	// fmt.Println("")
	// fmt.Println("grid", grid)

	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// fmt.Println(i, j)
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
			return 0
		}
		if visited[i][j] {
			return 0
		}

		visited[i][j] = true

		if grid[i][j] == 0 {
			return 0
		}

		return 1 + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j-1) + dfs(i, j+1)
	}

	var maxArea int
	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] != 1 {
				continue
			}

			area := dfs(i, j)
			// fmt.Println(area)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
