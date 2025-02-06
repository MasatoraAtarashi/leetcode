package main

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, len(board))
	columns := make([]map[byte]bool, len(board))
	boxes := make([]map[byte]bool, len(board))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if rows[i] == nil {
				rows[i] = make(map[byte]bool)
			}
			if columns[j] == nil {
				columns[j] = make(map[byte]bool)
			}

			boxIndex := (i/3)*3 + (j / 3)
			// fmt.Println("boxIndex", boxIndex)
			if boxes[boxIndex] == nil {
				boxes[boxIndex] = make(map[byte]bool)
			}

			num := board[i][j]
			// fmt.Println("num", num)
			if num == '.' {
				continue
			}

			if _, ok := rows[i][num]; ok {
				return false
			}
			if _, ok := columns[j][num]; ok {
				return false
			}
			if _, ok := boxes[boxIndex][num]; ok {
				return false
			}

			rows[i][num] = true
			columns[j][num] = true
			boxes[boxIndex][num] = true

			// fmt.Println("rows", rows)
			// fmt.Println("columns", columns)
			// fmt.Println("boxes", boxes)
		}
	}

	return true
}
