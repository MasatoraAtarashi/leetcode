package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	var maxReach int
	for i := range nums {
		if i > maxReach {
			return false
		}
		maxReach = maxint(i+nums[i], maxReach)
		if maxReach >= len(nums)-1 {
			fmt.Println("ok")
			return true
		}
	}
	return false
}
