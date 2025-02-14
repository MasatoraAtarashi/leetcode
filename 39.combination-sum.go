package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	// ソート
	sort.Ints(candidates)

	// バックトラック&枝刈り
	var result [][]int
	var current []int
	var backtrack func(idx, remaining int)
	backtrack = func(idx, remaining int) {
		// ぴったりが見つかった
		if remaining == 0 {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			fmt.Println("Yes!!", result)
			return
		}

		for i := idx; i <= len(candidates)-1; i++ {
			if candidates[i] > remaining {
				break
			}
			current = append(current, candidates[i])
			fmt.Println(current, i, remaining)
			backtrack(i, remaining-candidates[i])
			current = current[:len(current)-1]
		}
	}

	backtrack(0, target)
	return result
}
