package main

import "fmt"

func countSubstrings(s string) int {
	var isPali func(i, j int) bool
	isPali = func(i, j int) bool {
		fmt.Println("call")
		// 1文字
		if i == j {
			return true
		}

		// 2文字
		if i+1 == j {
			return s[i] == s[j]
		}

		// 3文字以上
		if s[i] != s[j] {
			return false
		}

		return isPali(i+1, j-1)
	}

	var count int
	for i := range len(s) {
		for j := i; j < len(s); j++ {
			if isPali(i, j) {
				count++
			}
		}
	}
	return count
}
