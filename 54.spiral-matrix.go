package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	fmt.Println("matrix", matrix)

	result := make([]int, 0, len(matrix)*len(matrix[0]))
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		fmt.Printf("top: %d, bottom: %d, left: %d, right: %d\n", top, bottom, left, right)
		// 一番上の左→右
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		fmt.Println("result", result)

		// 一番右の上から下
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		fmt.Println("result", result)

		// 一番下の右から左
		for j := right - 1; j >= left; j-- {
			if top < bottom {
				result = append(result, matrix[bottom][j])
			}
		}
		fmt.Println("result", result)

		// 一番左の下から上
		for i := bottom - 1; i > top; i-- {
			if top < bottom && left < right {
				result = append(result, matrix[i][left])
			}
		}
		fmt.Println("result", result)

		top++
		bottom--
		left++
		right--
	}

	return result
}
