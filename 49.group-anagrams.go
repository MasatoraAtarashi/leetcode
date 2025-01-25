package main

import "leetcode/pkg"

func groupAnagrams(strs []string) [][]string {
	resultss := [][]string{{strs[0]}}

	for i := 1; i < len(strs); i++ {
		anagram := false
		for j := range resultss {
			if pkg.IsAnagram(strs[i], resultss[j][0]) {
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
