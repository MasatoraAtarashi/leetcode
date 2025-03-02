package main

func canPartition(nums []int) bool {
	var total int
	for _, n := range nums {
		total += n
	}

	// 合計値が奇数なら絶対無理
	if total%2 != 0 {
		return false
	}

	target := total / 2

	dp := make([]bool, target+1)
	for i, n := range nums {
		// 空は必ず作れる
		if i == 0 {
			dp[i] = true
			continue
		}

		for j := target; j >= n; j-- {
			if dp[j-n] {
				dp[j] = true
				if j == target {
					return true
				}
			}
		}
	}

	return false
}
