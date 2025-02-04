package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		rs := []rune(str)
		sort.Slice(rs, func(i, j int) bool {
			return rs[i] < rs[j]
		})
		if _, ok := m[string(rs)]; ok {
			m[string(rs)] = append(m[string(rs)], str)
		} else {
			m[string(rs)] = []string{str}
		}
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}
