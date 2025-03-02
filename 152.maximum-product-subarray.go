package main

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max, maxIncludeNow, minIncludeNow := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// まず現在の数字を含めたときの最大値・最小値を計算
		if nums[i] == 0 {
			maxIncludeNow = 0
			minIncludeNow = 0
		} else {
			tmax := maxIncludeNow * nums[i]
			tmin := minIncludeNow * nums[i]
			// 負の数 * 負の数を考慮
			if tmax < tmin {
				tmax, tmin = tmin, tmax
			}
			maxIncludeNow = maxint(tmax, nums[i])
			minIncludeNow = minint(tmin, nums[i])
		}

		// どの部分配列よりも大きければ更新
		if maxIncludeNow > max {
			max = maxIncludeNow
		}

		// fmt.Println("max", max)
		// fmt.Println("maxIncludeNow", maxIncludeNow)
		// fmt.Println("minIncludeNow", minIncludeNow)
	}

	return max
}
