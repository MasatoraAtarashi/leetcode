package main

import "fmt"

func findMin(nums []int) int {
	i := 0
	j := len(nums) - 1
	min := 99999999999
	for i <= j {
		mid := (i + j) / 2
		fmt.Println("i: ", i, " j: ", j, " mid: ", mid)
		if nums[mid] < nums[j] {
			j = mid - 1
		} else {
			i = mid + 1
		}

		if nums[mid] < min {
			min = nums[mid]
		}
	}
	fmt.Println("min: ", min)
	return min
}
