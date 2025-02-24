package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	current := nums[0]
	max := current
	for i := 1; i < len(nums); i++ {
		// 引き継ぐか、ここから新しく始めるか判断
		if current+nums[i] < nums[i] {
			current = nums[i]
		} else {
			current += nums[i]
		}
		if current > max {
			max = current
		}
	}

	return max
}
