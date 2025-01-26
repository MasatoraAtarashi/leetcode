package main

func countBits(n int) []int {
	bs := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		bs[i] = hammingWeight(i)
	}
	return bs
}
