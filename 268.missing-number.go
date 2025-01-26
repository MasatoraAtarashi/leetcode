package main

func missingNumber(nums []int) int {
	result := 0
	for i, n := range nums {
		result ^= i ^ n
	}
	result ^= len(nums)

	return result
}
