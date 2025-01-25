package main

func IsAnagram(s string, t string) bool {
	runes := make([]int, 26)
	for _, ss := range s {
		runes[int(ss-'a')]++
	}

	for _, tt := range t {
		runes[int(tt)-'a']--
	}

	for _, r := range runes {
		if r != 0 {
			return false
		}
	}

	return true
}
