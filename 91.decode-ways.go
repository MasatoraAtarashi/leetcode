package main

import "strconv"

func numDecodings(s string) int {
	patterns := make([]int, len(s)+1)
	// 空文字想定
	patterns[0] = 1

	for i := range s {
		if i == 0 {
			if string(s[i]) == "0" {
				return 0
			}
			patterns[i+1] = 1
			continue
		}
		patterns[i+1] = 0

		// 単独で文字として使えるパターン
		if string(s[i]) != "0" {
			patterns[i+1] = patterns[i]
		}

		// 前の文字と合わせて使えるパターン

		if string(s[i-1]) != "0" {
			sn := string(s[i-1]) + string(s[i])
			nn, _ := strconv.Atoi(sn)
			if nn > 0 && nn <= 26 {
				patterns[i+1] += patterns[i-1]
			}
		}
	}

	return patterns[len(s)]
}
