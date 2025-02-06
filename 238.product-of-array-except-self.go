package main

func productExceptSelf(nums []int) []int {
	leftMulti := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			leftMulti[i] = 1
			continue
		}
		if i == 1 {
			leftMulti[i] = nums[i-1]
			continue
		}
		leftMulti[i] = leftMulti[i-1] * nums[i-1]
	}
	// fmt.Println("leftMulti", leftMulti)

	rightMulti := make([]int, len(nums))
	result := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightMulti[i] = 1
		} else if i == len(nums)-2 {
			rightMulti[i] = nums[i+1]
		} else {
			rightMulti[i] = rightMulti[i+1] * nums[i+1]
		}

		// fmt.Println("rightMulti", rightMulti)
		result[i] = leftMulti[i] * rightMulti[i]
		// fmt.Println("result", result)
	}
	return result
}
