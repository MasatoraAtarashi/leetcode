package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var current []int
	var backtrack3 func(idx int)
	backtrack3 = func(idx int) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for i := idx; i < len(nums)-1; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			current = append(current, nums[i])
			backtrack3(i + 1)
			if len(current) > 0 {
				current = current[:len(current)-1]
			}
		}
	}

	backtrack3(0)
	return result
}
