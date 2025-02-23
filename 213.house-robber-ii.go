package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxint(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return maxint(maxint(nums[0], nums[1]), maxint(nums[1], nums[2]))
	}

	// 最初の家込パターン
	prevprev := nums[0]
	prev := nums[1]
	var maxIn int
	var current int
	for i := 2; i < len(nums)-1; i++ {
		current = nums[i]
		tmpprev := prev
		prev = maxint(prevprev+current, prev)
		maxIn = prev
		prevprev = maxint(prevprev, tmpprev)
	}

	prevprev = nums[1]
	prev = nums[2]
	// 最初の家除外パターン
	var maxOut int
	for i := 3; i < len(nums); i++ {
		current = nums[i]
		tmpprev := prev
		prev = maxint(prevprev+current, prev)
		maxOut = prev
		prevprev = maxint(prevprev, tmpprev)
	}

	return maxint(maxIn, maxOut)
}
