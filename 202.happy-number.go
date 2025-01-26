package main

import "fmt"

func isHappy(n int) bool {
	checked := make(map[int]bool)
	for n != 1 {
		n = nextNumber(n)
		fmt.Println("next", n)
		if ok := checked[n]; ok {
			return false
		}

		checked[n] = true
	}

	return true
}

// 各桁の2乗の和
func nextNumber(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}
