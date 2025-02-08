package main

func searchMatrix(matrix [][]int, target int) bool {
	// 行単位のバイナリーサーチ
	i := 0
	j := len(matrix) - 1
	rowNumber := 0
	for i <= j {
		mid := (i + j) / 2
		// fmt.Printf("行バイナリーサーチ開始: i=%d, j=%d, mid=%d, 範囲=[%d, %d], target=%d\n",
		// 	i, j, mid, matrix[mid][0], matrix[mid][len(matrix[mid])-1], target)
		if target < matrix[mid][0] {
			j = mid - 1
			// fmt.Printf("行バイナリーサーチ: target(%d) < matrix[%d][0](%d) のため j を %d に更新\n",
			// 	target, mid, matrix[mid][0], j)
		} else if target <= matrix[mid][len(matrix[mid])-1] {
			rowNumber = mid
			// fmt.Printf("行バイナリーサーチ: target(%d) は matrix[%d] の範囲 [%d, %d] に含まれるので rowNumber を %d に設定\n",
			// 	target, mid, matrix[mid][0], matrix[mid][len(matrix[mid])-1], rowNumber)
			break
		} else {
			i = mid + 1
			// fmt.Printf("行バイナリーサーチ: target(%d) > matrix[%d][last](%d) のため i を %d に更新\n",
			// 	target, mid, matrix[mid][len(matrix[mid])-1], i)
		}
	}
	// fmt.Printf("行バイナリーサーチ終了: 選択された行 rowNumber=%d\n", rowNumber)

	// 列単位のバイナリーサーチ
	exists := false
	ii := 0
	jj := len(matrix[rowNumber]) - 1
	for ii <= jj {
		mid := (ii + jj) / 2
		// fmt.Printf("列バイナリーサーチ開始: ii=%d, jj=%d, mid=%d, 値=%d, target=%d\n",
		// ii, jj, mid, matrix[rowNumber][mid], target)
		if target == matrix[rowNumber][mid] {
			// fmt.Printf("列バイナリーサーチ: row %d, column %d で target(%d) を発見\n", rowNumber, mid, target)
			return true
		}
		if target < matrix[rowNumber][mid] {
			jj = mid - 1
			// fmt.Printf("列バイナリーサーチ: target(%d) < matrix[%d][%d](%d) のため jj を %d に更新\n",
			// 	target, rowNumber, mid, matrix[rowNumber][mid], jj)
		} else {
			ii = mid + 1
			// fmt.Printf("列バイナリーサーチ: target(%d) > matrix[%d][%d](%d) のため ii を %d に更新\n",
			// 	target, rowNumber, mid, matrix[rowNumber][mid], ii)
		}
	}
	// fmt.Printf("列バイナリーサーチ終了: 最終 ii=%d, jj=%d で target(%d) は見つかりませんでした\n", ii, jj, target)
	return exists
}
