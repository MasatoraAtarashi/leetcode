package main

func twoSum2(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		v := numbers[i] + numbers[j]
		if v == target {
			return []int{i + 1, j + 1}
		} else if v < target {
			i++
		} else {
			j--
		}
	}

	return []int{}
}
