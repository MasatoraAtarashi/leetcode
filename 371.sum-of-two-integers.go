package main

func getSum(a int, b int) int {
	// fmt.Printf("a (binary): %b\n", a)
	// fmt.Printf("b (binary): %b\n", b)

	// 以下を繰り上がりがなくなるまで繰り返す
	// 1. XORで繰り上がりなし部分計算
	// 2. AND→左シフトで繰り上がり部分計算
	// 3. 1に戻る
	noncarry := a ^ b

	carry := a & b << 1

	for carry != 0 {
		// fmt.Printf("nocarry (binary): %b\n", noncarry)
		// fmt.Printf("carry (binary): %b\n", carry)
		tmp := noncarry
		noncarry = noncarry ^ carry
		carry = tmp & carry << 1
	}

	return noncarry
}
