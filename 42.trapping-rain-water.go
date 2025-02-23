package main

func trap(height []int) int {
	var sum, leftMax, rightMax int
	l := 0
	r := len(height) - 1
	for l < r {
		if height[l] > leftMax {
			leftMax = height[l]
		}
		if height[r] > rightMax {
			rightMax = height[r]
		}
		if leftMax < rightMax {
			if l != 0 {
				sum += leftMax - height[l]
			}
			l++
		} else {
			if r != len(height)-1 {
				sum += rightMax - height[r]
			}
			r--
		}
	}

	return sum
}
