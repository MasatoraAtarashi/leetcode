package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	maxConsective := 0
	for key, _ := range m {
		if _, ok := m[key-1]; !ok {
			consective := 1
			current := key
			for {
				if _, ok := m[current+1]; ok {
					consective++
					current++
				} else {
					break
				}
			}

			if maxConsective < consective {
				maxConsective = consective
			}
		}
	}

	return maxConsective
}
