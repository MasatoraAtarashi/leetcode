package main

func longestPalindrome(s string) string {
	var max, current string

	var check func(c1, c2 int)
	check = func(c1, c2 int) {
		if c1 < 0 || c2 >= len(s) {
			return
		}
		if s[c1] != s[c2] {
			return
		}
		current = s[c1 : c2+1]
		if c1-1 >= 0 && c2+1 < len(s) {
			check(c1-1, c2+1)
		}
	}

	for i := range s {
		// 奇数回文
		check(i, i)
		if len(current) > len(max) {
			max = current
		}

		// 偶数回文
		check(i, i+1)

		if len(current) > len(max) {
			max = current
		}
	}

	return max
}
