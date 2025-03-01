package main

import "fmt"

func rotate(matrix [][]int) {
	// fmt.Println("matrix", matrix)

	// 転置
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if i < j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		}
	}
	// fmt.Println("transposed", matrix)

	// 左右反転
	for i := range len(matrix) {
		for j := range len(matrix[i]) / 2 {
			fmt.Println(i, j)
			matrix[i][j], matrix[i][len(matrix)-1-j] = matrix[i][len(matrix)-1-j], matrix[i][j]
		}
	}
	// fmt.Println("horizontal fliped", matrix)
}
