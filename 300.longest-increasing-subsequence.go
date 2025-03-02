package main

func lengthOfLIS(nums []int) int {
	lislist := make([][]int, len(nums))
	for i := range nums {
		lislist[i] = make([]int, len(nums))
		for j := range nums {
			lislist[i][j] = -1
		}
	}

	var lis func(idx, prev int) int
	lis = func(idx, prev int) int {
		// fmt.Println(idx, prev)
		if idx >= len(nums) {
			return 0
		}

		if prev != -1 {
			if lislist[idx][prev] != -1 {
				// fmt.Println("メモ活用")
				return lislist[idx][prev]
			}
		}

		// fmt.Println("Notメモ")

		// 現在の値を選べる
		if prev == -1 {
			l := maxint(1+lis(idx+1, idx), lis(idx+1, prev))
			// fmt.Printf("idx: %d, prev: %d, l: %d\n", idx, prev, l)
			return l
		}

		if nums[idx] > nums[prev] {
			l := maxint(1+lis(idx+1, idx), lis(idx+1, prev))
			lislist[idx][prev] = l
			// fmt.Printf("idx: %d, prev: %d, l: %d\n", idx, prev, l)
			return l
		}

		// 選べない
		l := lis(idx+1, prev)
		lislist[idx][prev] = l
		// fmt.Printf("idx: %d, prev: %d, l: %d\n", idx, prev, l)
		return l
	}

	return lis(0, -1)
}
