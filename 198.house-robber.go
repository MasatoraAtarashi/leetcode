package main

func _rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxint(nums[0], nums[1])
	}
	prevprev := nums[0]
	prev := nums[1]
	var current int
	for i := 2; i < len(nums); i++ {
		current = nums[i]
		tmpprev := prev
		prev = maxint(prevprev+current, prev)
		prevprev = maxint(prevprev, tmpprev)
	}
	return prev
}

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}
