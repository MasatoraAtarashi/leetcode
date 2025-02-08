package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j > (i+1) && nums[j] == nums[j-1] {
				j++
				continue
			}

			val := nums[i] + nums[j] + nums[k]
			fmt.Printf("val: %d, i: %d, j: %d, k: %d\n", val, nums[i], nums[j], nums[k])
			if val == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
			} else if val < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result
}
