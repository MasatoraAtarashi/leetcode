package main

func rob(nums []int) int {
	var robrob func(nums []int) int
	robrob = func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		if len(nums) == 1 {
			return nums[0]
		}
		if len(nums) == 2 {
			return maxx(nums[0], nums[1])
		}

		if max[len(nums)-1] != 0 {
			return max[len(nums)-1]
		}

		p1 := robrob(nums[:len(nums)-2]) + nums[len(nums)-1]
		p2 := robrob(nums[:len(nums)-1])
		if p1 < p2 {
			max[len(nums)-1] = p2
			return p2
		}
		max[len(nums)-1] = p1
		return p1
	}

	return robrob(nums)
}

func maxint(a, b int) int {
	if a > b {
		return a
	}
	return b
}
