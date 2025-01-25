package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	i := 0
	j := len(s) - 1
	for {
		if i == j {
			return true
		}
		if i > j {
			return true
		}

		for {
			if i >= len(s) {
				return true
			}

			if IsAlphanumeric(rune(s[i])) {
				break
			} else {
				i++
			}
		}

		for {
			if j < 0 {
				return true
			}

			if IsAlphanumeric(rune(s[j])) {
				break
			} else {
				j--
			}
		}

		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}

		i++
		j--
	}
}
