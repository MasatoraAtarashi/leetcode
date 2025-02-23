package main

func permute(nums []int) [][]int {
	var result [][]int
	current := make([]int, 0, len(nums))
	var f func(nums []int)
	f = func(ns []int) {
		if len(ns) == 0 {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}

		for i := range len(ns) {
			current = append(current, ns[i])
			f(RemoveIndex(ns, i))

			// backtrack
			current = current[:len(current)-1]
		}
	}

	f(nums)
	return result
}
