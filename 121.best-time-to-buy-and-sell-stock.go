package main

func maxProfit(prices []int) int {
	var maxProfit, minProfit int
	minProfit = 999999999999999999
	for i := 0; i < len(prices); i++ {
		if prices[i] < minProfit {
			minProfit = prices[i]
		}

		if prices[i]-minProfit > maxProfit {
			maxProfit = prices[i] - minProfit
		}
	}

	if maxProfit < 0 {
		return 0
	}
	return maxProfit
}
