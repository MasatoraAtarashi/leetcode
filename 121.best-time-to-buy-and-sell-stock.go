package main

func maxProfit(prices []int) int {
	var maxProfit int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}

	if maxProfit < 0 {
		return 0
	}
	return maxProfit
}
