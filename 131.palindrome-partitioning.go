package main

import "fmt"

func partition(s string) [][]string {
	var result [][]string
	current := make([]string, 0, len(s))
	var ff func(start int)
	ff = func(start int) {
		// fmt.Println("start", start)
		if start == len(s) {
			tmp := make([]string, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			// fmt.Println("result added", result)
			return
		}

		for i := start + 1; i <= len(s); i++ {
			fmt.Println("i", i)
			now := string(s[start:i])
			// fmt.Println("now", now)
			if !isPalindrome(now) {
				continue
			}

			current = append(current, now)
			// fmt.Println("current", current)

			ff(i)
			// fmt.Println("backtrack")
			current = current[:len(current)-1]
		}
	}

	ff(0)
	return result
}
