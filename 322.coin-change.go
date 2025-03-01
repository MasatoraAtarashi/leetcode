package main

func coinChange(coins []int, amount int) int {
	mins := make([]int, amount+1)
	for i := range amount + 1 {
		// 理論上最大値
		mins[i] = amount + 1
	}
	mins[0] = 0

	for _, coin := range coins {
		// fmt.Println(coin)
		for i := range amount + 1 {
			if i >= coin {
				mins[i] = minint(mins[i], mins[i-coin]+1)
			}
		}
		// fmt.Println(mins)
	}

	// 与えられたコインでは払えない
	if mins[amount] > amount {
		return -1
	}
	return mins[amount]
}
