package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	var current []int
	var backtrack2 func(idx, remaining int)
	backtrack2 = func(idx, remaining int) {
		if remaining == 0 {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if candidates[i] > remaining {
				fmt.Println("Over", i, idx)
				break
			}
			if i > idx && candidates[i] == candidates[i-1] {
				fmt.Println("Skip!")
				continue
			}
			current = append(current, candidates[i])
			fmt.Println(current, i, remaining)
			backtrack2(i, remaining-candidates[i])
			current = current[:len(current)-1]
		}
	}

	backtrack2(0, target)
	return result
}
