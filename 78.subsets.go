package main

func subsets(nums []int) [][]int {
	var result [][]int
	var current []int
	var backtrack func(idx int)
	backtrack = func(idx int) {
		// 追加
		tmp := make([]int, len(current))
		copy(tmp, current)
		result = append(result, tmp)

		for i := idx; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i + 1)
		}
		if len(current) > 0 {
			current = current[:len(current)-1]
		}
	}

	backtrack(0)
	return result
}
