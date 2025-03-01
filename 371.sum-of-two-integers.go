package main

func getSum(a int, b int) int {
	// 以下を繰り上がりがなくなるまで繰り返す
	// 1. XORで繰り上がりなし部分計算
	// 2. AND→左シフトで繰り上がり部分計算
	// 3. 1に戻る
	noncarry := a ^ b
	carry := a & b 
	carry = carr	
	return 0
}
