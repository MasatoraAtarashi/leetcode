package main

func plusOne(digits []int) []int {
	allNine := true
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			allNine = false
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	if allNine {
		var result = []int{1}
		for i := 0; i < len(digits); i++ {
			result = append(result, 0)
		}
		return result
	}

	return digits
}
