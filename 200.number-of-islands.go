package main

func numIslands(grid [][]byte) int {
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row >= len(grid) || row < 0 || col >= len(grid[row]) || col < 0 {
			return
		}
		if grid[row][col] != '1' {
			return
		}
		// 訪問済み
		grid[row][col] = '0'
		dfs(row+1, col)
		dfs(row, col+1)
		dfs(row-1, col)
		dfs(row, col-1)
	}

	var islands int
	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == '1' {
				dfs(i, j)
				islands++
			}
		}
	}

	return islands
}
