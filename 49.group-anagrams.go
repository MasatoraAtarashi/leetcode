package main

func groupAnagrams(strs []string) [][]string {
	resultss := [][]string{{strs[0]}}

	for i := 1; i < len(strs); i++ {
		anagram := false
		for j := range resultss {
			if isAnagram(strs[i], resultss[j][0]) {
				resultss[j] = append(resultss[j], strs[i])
				anagram = true
				continue
			}
		}

		if !anagram {
			resultss = append(resultss, []string{strs[i]})
		}
	}

	return resultss
}

func isAnagram(s string, t string) bool {
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
