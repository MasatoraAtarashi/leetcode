package main

func reverse(x int) int {
	max := 1<<31 - 1
	min := -1 << 31

	var result int
	for x != 0 {
		t := x % 10
		result = result*10 + t
		if result > max || result < min {
			return 0
		}
		x = x / 10
	}

	return result
}
