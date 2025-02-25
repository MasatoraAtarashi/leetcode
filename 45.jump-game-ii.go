package main

func jump(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	var jumpCnt, currentEnd, farthest int
	for i := range nums {
		if currentEnd >= len(nums)-1 {
			break
		}
		farthest = maxint(farthest, i+nums[i])
		if currentEnd == i {
			jumpCnt++
			currentEnd = farthest

		}
	}
	return jumpCnt
}
