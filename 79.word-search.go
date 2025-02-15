package main

const DONE = '#'

func exist(board [][]byte, word string) bool {
	var search func(row, col, wordIndex int) bool
	search = func(row, col, wordIndex int) bool {
		// fmt.Println(row, col, wordIndex)
		if wordIndex == len(word) {
			return true
		}

		if row >= len(board) || row < 0 || col >= len(board[row]) || col < 0 {
			return false
		}
		// fmt.Println(string(board[row][col]), string(word[wordIndex]))
		if board[row][col] == DONE {
			return false
		}
		if board[row][col] != word[wordIndex] {
			return false
		}

		tmp := board[row][col]
		board[row][col] = DONE
		result := search(row+1, col, wordIndex+1) || search(row, col+1, wordIndex+1) || search(row-1, col, wordIndex+1) || search(row, col-1, wordIndex+1)

		board[row][col] = tmp

		return result
	}

	for i := range len(board) {
		for j := range len(board[i]) {
			if search(i, j, 0) {
				return true
			}
		}
	}

	return false
}
