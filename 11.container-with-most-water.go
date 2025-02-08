package main

func maxArea(height []int) int {
	maxHeight := 0
	i := 0
	j := len(height) - 1
	for i < j {
		width := j - i
		h := height[i]
		if height[j] < h {
			h = height[j]
		}

		area := width * h
		if maxHeight < area {
			maxHeight = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxHeight
}
