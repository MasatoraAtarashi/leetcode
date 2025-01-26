package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	costsToClime := make([]int, len(cost)+1)
	costsToClime[0] = 0
	costsToClime[1] = 0
	for i := 2; i <= len(cost); i++ {
		fmt.Println(i)
		costsToClime[i] = min(cost[i-1]+costsToClime[i-1], cost[i-2]+costsToClime[i-2])
		fmt.Println(costsToClime)
	}

	return costsToClime[len(cost)]
}

func main() {
	minCostClimbingStairs([]int{10, 15, 20})
}
