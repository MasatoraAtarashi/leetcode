package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	freqmap := make(map[int]int)
	for _, num := range nums {
		freqmap[num]++
	}

	keys := make([]int, 0, len(freqmap))
	for key := range freqmap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return freqmap[keys[i]] > freqmap[keys[j]]
	})

	return keys[:k]
}
