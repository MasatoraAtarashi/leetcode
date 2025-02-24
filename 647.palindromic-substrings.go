package main

import "fmt"

func countSubstrings(s string) int {
	// 0: default, 1: palindromic, -1: not
	memo := make([][]int, len(s))

	var isPali func(i, j int) bool
	isPali = func(i, j int) bool {
		fmt.Println("call")
		fmt.Println(i, j)

		if i < 0 || j >= len(s) || i > j {
			return false
		}

		if memo[i] == nil {
			memo[i] = make([]int, len(s))
		}

		if memo[i][j] != 0 {
			if memo[i][j] == 1 {
				return true
			} else {
				return false
			}
		}

		// 1文字
		if i == j {
			memo[i][j] = 1
			return true
		}

		// 2文字
		if i+1 == j {
			if s[i] == s[j] {
				memo[i][j] = 1
				return true
			} else {
				memo[i][j] = -1
				return false
			}
		}

		// 3文字以上
		if s[i] != s[j] {
			memo[i][j] = -1
			return false
		}

		if isPali(i+1, j-1) {
			memo[i][j] = 1
			return true
		}

		memo[i][j] = -1
		return false
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
