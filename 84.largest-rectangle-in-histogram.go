package main

import "fmt"

func largestRectangleArea(heights []int) int {
	leftLimits := make([]int, len(heights))
	leftStack := Stack[int]{}
	for i := range heights {
		if leftStack.Empty() {
			leftLimits[i] = -1
			leftStack.Push(i)
			continue
		}
		if heights[i] > heights[leftStack.Top()] {
			leftLimits[i] = leftStack.Top()
			leftStack.Push(i)
			continue
		}

		for !leftStack.Empty() && heights[i] <= heights[leftStack.Top()] {
			leftStack.Pop()
		}

		if leftStack.Empty() {
			leftStack.Push(i)
			leftLimits[i] = -1
		} else {
			leftLimits[i] = leftStack.Top()
			leftStack.Push(i)
		}
	}
	// fmt.Println("leftLimits", leftLimits)

	rightLimits := make([]int, len(heights))
	rightStack := Stack[int]{}
	for i := len(heights) - 1; i >= 0; i-- {
		if rightStack.Empty() {
			rightLimits[i] = -1
			rightStack.Push(i)
			continue
		}

		if heights[i] > heights[rightStack.Top()] {
			rightLimits[i] = rightStack.Top()
			rightStack.Push(i)
			continue
		}

		for !rightStack.Empty() && heights[i] <= heights[rightStack.Top()] {
			rightStack.Pop()
		}

		if rightStack.Empty() {
			rightStack.Push(i)
			rightLimits[i] = -1
		} else {
			rightLimits[i] = rightStack.Top()
			rightStack.Push(i)
		}
	}
	// fmt.Println("rightLimits", rightLimits)

	var max int
	for i := range heights {
		var l int
		if leftLimits[i] == -1 {
			l = 0
		} else {
			l = leftLimits[i] + 1
		}

		r := rightLimits[i]
		if r == -1 {
			r = len(heights)
		}

		area := heights[i] * (r - l)
		if area > max {
			fmt.Println("max", area)
			max = area
		}
	}

	return max
}
