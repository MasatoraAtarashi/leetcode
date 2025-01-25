package main

import "unicode"

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		for i < j && !IsAlphanumeric(rune(s[i])) {
			i++
		}

		for i < j && !IsAlphanumeric(rune(s[j])) {
			j--
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}
	return true
}
